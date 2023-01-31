package com.bookservice.bookservice.query.controller;

import com.bookservice.bookservice.command.controller.BookCommandController;
import com.bookservice.bookservice.query.model.BookResponseModel;
import com.bookservice.bookservice.query.queries.GetBookQuery;
import org.axonframework.messaging.responsetypes.ResponseType;
import org.axonframework.messaging.responsetypes.ResponseTypes;
import org.axonframework.queryhandling.QueryGateway;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/books")
public class BookQueryController {
    private QueryGateway queryGateway;

    public BookQueryController(QueryGateway queryGateway) {
        this.queryGateway = queryGateway;
    }

    @GetMapping("/{bookId}")
    public ResponseEntity<BookResponseModel> getBookDetail(@PathVariable(name = "bookId") String bookId) {
        final GetBookQuery getBookQuery = new GetBookQuery();
        getBookQuery.setBookId(bookId);
        final BookResponseModel bookResponseModel = queryGateway.query(getBookQuery,
                ResponseTypes.instanceOf(BookResponseModel.class))
                .join();
        return ResponseEntity.ok(bookResponseModel);
    }
}
