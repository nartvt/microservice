package com.bookservice.bookservice.command.controller;

import com.bookservice.bookservice.command.command.CreateBookCommand;
import com.bookservice.bookservice.command.command.DeleteBookCommand;
import com.bookservice.bookservice.command.command.UpdateBookCommand;
import com.bookservice.bookservice.command.model.request.BookRequestModel;
import org.axonframework.commandhandling.gateway.CommandGateway;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.UUID;

@RestController
@RequestMapping("/api/v1/books")
public class BookCommandController {
    private CommandGateway commandGateway;

    public BookCommandController(CommandGateway commandGateway) {
        this.commandGateway = commandGateway;
    }

    @PostMapping
    public ResponseEntity<String> addBook(@RequestBody BookRequestModel model) {
        final CreateBookCommand command = new CreateBookCommand(
                UUID.randomUUID().toString(),
                model.getName(),
                model.getAuthor(),
                model.isReady());
        commandGateway.sendAndWait(command);
        return ResponseEntity.ok("created Book");
    }

    @PutMapping
    public ResponseEntity<String> updateBook(@RequestBody BookRequestModel model) {
        final UpdateBookCommand command = new UpdateBookCommand(
                model.getBookId(),
                model.getName(),
                model.getAuthor(),
                model.isReady());
        commandGateway.sendAndWait(command);
        return ResponseEntity.ok("updated Book");
    }

    @DeleteMapping("/{bookId}")
    public ResponseEntity<String> deleteBook(@PathVariable("bookId") String bookId) {
        final DeleteBookCommand command = new DeleteBookCommand(bookId);
        commandGateway.sendAndWait(command);
        return ResponseEntity.ok("deleted Book");
    }
}
