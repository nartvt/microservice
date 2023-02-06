package com.bookservice.bookservice.command.aggregate;

import com.bookservice.bookservice.command.command.CreateBookCommand;
import com.bookservice.bookservice.command.command.DeleteBookCommand;
import com.bookservice.bookservice.command.command.UpdateBookCommand;
import com.bookservice.bookservice.command.event.CreateBookEvent;
import com.bookservice.bookservice.command.event.DeleteBookEvent;
import com.bookservice.bookservice.command.event.UpdateBookEvent;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.axonframework.commandhandling.CommandHandler;
import org.axonframework.eventsourcing.EventSourcingHandler;
import org.axonframework.modelling.command.AggregateIdentifier;
import org.axonframework.modelling.command.AggregateLifecycle;
import org.axonframework.spring.stereotype.Aggregate;
import org.springframework.beans.BeanUtils;

@Aggregate
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class BookAggregate {
    @AggregateIdentifier
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;

    @CommandHandler
    public BookAggregate(CreateBookCommand command) {
        final CreateBookEvent event = new CreateBookEvent();
        BeanUtils.copyProperties(command, event);
        AggregateLifecycle.apply(event);
    }

    @CommandHandler
    public void handle(UpdateBookCommand command) {
        final UpdateBookEvent event = new UpdateBookEvent();
        event.setBookId(command.getBookId());
        event.setName(command.getName());
        event.setAuthor(command.getAuthor());
        event.setReady(command.isReady());
        AggregateLifecycle.apply(event);
    }

    @CommandHandler
    public void handle(DeleteBookCommand command) {
        final DeleteBookEvent event = new DeleteBookEvent();
        BeanUtils.copyProperties(command, event);
        AggregateLifecycle.apply(event);
    }

    @EventSourcingHandler
    public void on(CreateBookEvent event) {
        this.bookId = event.getBookId();
        this.name = event.getName();
        this.isReady = event.isReady();
        this.author = event.getAuthor();
    }

    @EventSourcingHandler
    public void on(UpdateBookEvent event) {
        this.bookId = event.getBookId();
        this.name = event.getName();
        this.isReady = event.isReady();
        this.author = event.getAuthor();
    }

    @EventSourcingHandler
    public void on(DeleteBookEvent event) {
        this.bookId = event.getBookId();
    }
}
