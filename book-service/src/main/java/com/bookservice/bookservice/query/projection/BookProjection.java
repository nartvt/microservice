package com.bookservice.bookservice.query.projection;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;
import com.bookservice.bookservice.query.model.BookResponseModel;
import com.bookservice.bookservice.query.queries.GetBookQuery;
import org.axonframework.queryhandling.QueryHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class BookProjection {

    @Autowired
    private IBookRepository bookRepository;

    @QueryHandler
    public BookResponseModel handle(GetBookQuery getBookQuery){
        BookResponseModel model =  new BookResponseModel();
        Book book = bookRepository.getReferenceById(getBookQuery.getBookId());
        BeanUtils.copyProperties(book,model);
        return model;
    }
}
