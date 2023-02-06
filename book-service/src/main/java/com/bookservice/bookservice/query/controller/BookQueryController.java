package com.bookservice.bookservice.query.controller;

import com.bookservice.bookservice.query.model.BookResponse;
import com.bookservice.bookservice.query.model.ResponseModel;
import com.bookservice.bookservice.query.queries.GetBookQuery;
import com.bookservice.bookservice.query.queries.GetBooksQueries;
import org.axonframework.messaging.responsetypes.ResponseTypes;
import org.axonframework.queryhandling.QueryGateway;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/v1/books")
public class BookQueryController {
    private QueryGateway queryGateway;

    public BookQueryController(QueryGateway queryGateway) {
        this.queryGateway = queryGateway;
    }

    @GetMapping("/{bookId}")
    public ResponseEntity<ResponseModel> getBookDetail(@PathVariable(name = "bookId") String bookId) {
        final GetBookQuery getBookQuery = new GetBookQuery();
        getBookQuery.setBookId(bookId);
        final BookResponse responseModel = queryGateway.query(getBookQuery,
                        ResponseTypes.instanceOf(BookResponse.class))
                .join();
        return ResponseEntity.ok(new ResponseModel(responseModel));
    }

    @GetMapping
    public ResponseEntity<ResponseModel> getBooks() {
        final GetBooksQueries getBookQuery = new GetBooksQueries();
        final List<BookResponse> responseModels = queryGateway.query(getBookQuery,
                        ResponseTypes.multipleInstancesOf(BookResponse.class))
                .join();
        return ResponseEntity.ok(new ResponseModel(responseModels));
    }
}
