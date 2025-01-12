import typing
import strawberry
from strawberry.fastapi import GraphQLRouter


@strawberry.type
class Book:
    title: str
    author: str

def get_books():
    return [
        Book(
            title='The Great Gatsby',
            author='F. Scott Fitzgerald',
        ),
    ]

@strawberry.type
class Query:
    books: typing.List[Book] = strawberry.field(resolver=get_books)

@strawberry.type
class Mutation:
    @strawberry.mutation
    def add_book(self, title: str, author: str) -> Book:
        print(f'Adding {title} by {author}')

        return Book(title=title, author=author)

schema = strawberry.Schema(query=Query, mutation=Mutation)

graphql_app = GraphQLRouter(schema)
