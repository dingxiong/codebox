from dataclasses import dataclass
import sys


@dataclass(frozen=True)
class Token:
    pos_from: int
    pos_to: int


@dataclass(frozen=True)
class Terminal(Token):
    body: str


@dataclass(frozen=True)
class Operator(Token):
    body: str


class Expression:
    pass


@dataclass(frozen=True)
class ExpressionLeaf(Expression):
    body: str

    """
    def __repr__(self) -> str:
        return self.body

    __str__ = __repr__
    """


@dataclass(frozen=True)
class ExpressionNonLeaf(Expression):
    op: Operator
    children: list[Expression]

    def repr(self, indent: int = 0) -> str:
        ws = " " * indent
        ret = f"{ws}op:{self.op.body}\n"
        for child in self.children:
            match child:
                case ExpressionLeaf() as leaf:
                    ret += f"{ws}  {leaf}\n"
                case ExpressionNonLeaf() as non_leaf:
                    ret += f"{non_leaf.repr(indent+2)}"
        return ret

    __repr__ = repr
    __str__ = repr


class Lexer:
    def __init__(self, code: str) -> None:
        self.code = code
        self.next_idx: int = 0
        self.current_token: Token | None = None

        # a stack to detach unbalanced delimiters such as parenthesis and bracket.
        self.delimiter_stack = []

    def _skip_white_spaces(self) -> None:
        while self.next_idx < len(self.code) and self.code[self.next_idx] in [
            " ",
            "\t",
        ]:
            self.next_idx += 1

    def has_next(self) -> bool:
        self._skip_white_spaces()
        return self.next_idx < len(self.code)

    def next(self) -> Token:
        self.peek()
        ret = self.current_token
        assert ret
        self.next_idx = ret.pos_to
        self.current_token = None  # make sure it is consumed.
        return ret

    def peek(self) -> Token:
        if self.current_token:
            return self.current_token
        self._skip_white_spaces()
        assert self.has_next()
        c = self.code[self.next_idx]
        match c:
            case "*" | "/":
                self.current_token = Operator(self.next_idx, self.next_idx + 1, str(c))
            case "(" | "[":
                self.current_token = Operator(self.next_idx, self.next_idx + 1, str(c))
                self.delimiter_stack.append(c)
            case ")" | "]":
                self.current_token = Operator(self.next_idx, self.next_idx + 1, str(c))
                assert self.delimiter_stack and self.delimiter_stack[-1] == (
                    "(" if c == ")" else "["
                ), f"Unbanced delimiter found {c}"
                self.delimiter_stack.pop()
            case "+":
                if (
                    self.next_idx + 1 < len(self.code)
                    and self.code[self.next_idx + 1] == "+"
                ):
                    self.current_token = Operator(
                        self.next_idx, self.next_idx + 2, "++"
                    )
                else:
                    self.current_token = Operator(
                        self.next_idx, self.next_idx + 1, str(c)
                    )
            case "-":
                if (
                    self.next_idx + 1 < len(self.code)
                    and self.code[self.next_idx + 1] == "-"
                ):
                    self.current_token = Operator(
                        self.next_idx, self.next_idx + 2, "--"
                    )
                else:
                    self.current_token = Operator(
                        self.next_idx, self.next_idx + 1, str(c)
                    )
            case c if "a" <= c <= "z" or "A" <= c <= "Z" or c == "_":
                j = 1
                while self.next_idx + j < len(self.code):
                    next_c = self.code[self.next_idx + j]
                    if (
                        "a" <= next_c <= "z"
                        or "A" <= next_c <= "Z"
                        or next_c == "_"
                        or "0" <= next_c <= "9"
                    ):
                        j += 1
                    else:
                        break
                self.current_token = Terminal(
                    self.next_idx,
                    self.next_idx + j,
                    self.code[self.next_idx : self.next_idx + j],
                )
            case c if "0" <= c <= "9":
                j = 1
                while self.next_idx + j < len(self.code):
                    next_c = self.code[self.next_idx + j]
                    if "0" <= next_c <= "9":
                        j += 1
                    else:
                        break
                self.current_token = Terminal(
                    self.next_idx,
                    self.next_idx + j,
                    self.code[self.next_idx : self.next_idx + j],
                )
            case _:
                raise ValueError(f"Invalid input at {self.next_idx}: {c}")

        return self.current_token

    def print_state(self):
        """For debug purpose"""
        print(self.code)
        print(f"{'~' * self.next_idx}^")


"""
Take the pseudocode from https://www.crockford.com/javascript/tdop/tdop.html

var expression = function (rbp) {
  var left;
  var t = token;
  advance();
  left = t.nud();
  while (rbp < token.lbp) {
    t = token;
    advance();
    left = t.led(left);
  }
  return left;
};

Binding power used in the below code.

infix
+, -: 5, 6
*, /: 7, 8
[, (: 1, 2

prefix
+, -  : _, 9
++, --: -, 10 

postfix:
++, --: 11, _ 
"""


def parse(code: str) -> Expression | None:
    lexer = Lexer(code)
    if not lexer.has_next():
        return None
    return parse_pratt(lexer, 0)


def parse_pratt(lexer: Lexer, rbp: int) -> Expression:
    if not lexer.has_next():
        raise ValueError("EOF")
    token = lexer.next()
    left: Expression | None = None
    match token:
        case Terminal(_, _, body):
            left = ExpressionLeaf(body)
        case Operator(_, _, body):
            match body:
                case "+" | "-":
                    right = parse_pratt(lexer, 9)
                    left = ExpressionNonLeaf(op=token, children=[right])
                case "++" | "--":
                    right = parse_pratt(lexer, 10)
                    left = ExpressionNonLeaf(op=token, children=[right])
                case "(":
                    right = parse_pratt(lexer, 0)
                    next = lexer.peek()
                    assert isinstance(next, Operator) and next.body == ")"
                    lexer.next()
                    left = ExpressionNonLeaf(op=token, children=[right])
                case _:
                    raise ValueError()
    assert left
    while True:
        # if EOF, then we should immediately return left.
        if not lexer.has_next():
            break

        # We should be peek here and only consume it when `rbp < token.lbp` holds.
        token = lexer.peek()

        assert isinstance(token, Operator), "The next token must be an operator"
        match token.body:
            case ")":
                break
            case "+" | "-":
                if rbp < 5:
                    lexer.next()
                    right = parse_pratt(lexer, 6)
                    left = ExpressionNonLeaf(op=token, children=[left, right])
                else:
                    break
            case "*" | "/":
                if rbp < 7:
                    lexer.next()
                    right = parse_pratt(lexer, 8)
                    left = ExpressionNonLeaf(op=token, children=[left, right])
                else:
                    break
            case "++" | "--":
                if rbp < 11:
                    lexer.next()
                    left = ExpressionNonLeaf(op=token, children=[left])
                else:
                    break
            case _:
                raise ValueError(f"unexpected operator {token}")

    return left


if __name__ == "__main__":
    test = sys.argv[1] if len(sys.argv) > 1 else "all"

    if test in ["t1", "all"]:
        print("==== test 1 ====")
        e = parse("a + b")
        print(e)

    if test in ["t2", "all"]:
        print("==== test 2 ====")
        e = parse("a + b * c")
        print(e)

    if test in ["t3", "all"]:
        print("==== test 3 ====")
        e = parse("(a + b) * c")
        print(e)

    if test in ["t4", "all"]:
        print("==== test 4 ====")
        e = parse("(a - ++b) * c")
        print(e)

    if test in ["t5", "all"]:
        print("==== test 5 ====")
        e = parse("(a---b) * c")
        print(e)

    if test in ["t6", "all"]:
        print("==== test 6 ====")
        try:
            e = parse("a---b) * c")
            print(e)
        except AssertionError as err:
            print(err)
