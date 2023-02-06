package com.bookservice.bookservice.query.projection;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.axonframework.queryhandling.QueryHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Component;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;
import com.bookservice.bookservice.query.model.BookResponse;
import com.bookservice.bookservice.query.queries.GetBookQuery;
import com.bookservice.bookservice.query.queries.GetBooksQueries;

@Component
public class BookProjection {

    private IBookRepository bookRepository;

    public BookProjection(IBookRepository bookRepository) {
        this.bookRepository = bookRepository;
    }

    @QueryHandler
    public BookResponse handle(GetBookQuery getBookQuery) {
        final BookResponse model = new BookResponse();
        final Book book = bookRepository.getReferenceById(getBookQuery.getBookId());
        BeanUtils.copyProperties(book, model);
        return model;
    }

    @QueryHandler
    public List<BookResponse> handle(GetBooksQueries booksQueries) {
        final List<Book> books = bookRepository.findAll();
        if (books.isEmpty()) {
            return Collections.emptyList();
        }
        final List<BookResponse> models = new ArrayList<>();
        books.forEach(book -> {
            final BookResponse model = new BookResponse();
            BeanUtils.copyProperties(book, model);
            models.add(model);
        });
        return models;
    }
}
