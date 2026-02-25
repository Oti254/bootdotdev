class Book:
    def __init__(self, title, author):
        self.title = title
        self.author = author

class Library:
    def __init__(self, name):
        self.name = name
        self.books = []

    def add_book(self, book):
        self.books.append(book)

    def remove_book(self, book):
        new_books = []
        for lib_book in self.books:
            if (lib_book.title != book.title) or (lib_book.author != book.author):
                new_books.append(lib_book)
        self.books = new_books

    def search_books(self, search_string):
        results = []
        for book in self.books:
            if (search_string in book.title.lower()) or (search_string in book.author.lower()):
                results.append(book)
        return results
