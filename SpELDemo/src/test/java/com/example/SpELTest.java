package com.example;

import org.junit.jupiter.api.Test;
import org.springframework.expression.Expression;
import org.springframework.expression.ExpressionParser;
import org.springframework.expression.spel.standard.SpelExpressionParser;
import org.springframework.expression.spel.support.StandardEvaluationContext;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class SpELTest {

    @Test
    public void testBasicExpression() {
        ExpressionParser parser = new SpelExpressionParser();
        String expression = "'Hello, ' + 'World!'";
        String result = parser.parseExpression(expression).getValue(String.class);
        assertEquals("Hello, World!", result);
    }

    @Test
    public void testConditionExpression() {
        ExpressionParser parser = new SpelExpressionParser();
        String expr = "'a' == 'b' ? 'option1' : 'option2'";
        String result = parser.parseExpression(expr).getValue(String.class);
        assertEquals("option2", result);
    }

    @Test
    public void testWithContext() {
        ExpressionParser parser = new SpelExpressionParser();
        StandardEvaluationContext context = new StandardEvaluationContext();
        context.setVariable("greeting", "Hello");
        String expression = "#greeting + ', ' + 'SpEL!'";
        String result = parser.parseExpression(expression).getValue(context, String.class);
        assertEquals("Hello, SpEL!", result);
    }

    class FooA {
        public String name;
        public FooA(){
            name = "ok";
        }
    }

    @Test
    public void testWithContextRoot() {
        FooA c = new FooA();
        ExpressionParser parser = new SpelExpressionParser();
        Expression exp = parser.parseExpression("name == 'ok'");
        boolean result = exp.getValue(c, Boolean.class);
        assertTrue(result);
    }
}
