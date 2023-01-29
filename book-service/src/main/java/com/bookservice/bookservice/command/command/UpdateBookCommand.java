package com.bookservice.bookservice.command.command;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import org.axonframework.modelling.command.TargetAggregateIdentifier;

@Getter
@Setter
public class UpdateBookCommand {
    @TargetAggregateIdentifier
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;
    public UpdateBookCommand(String bookId,String name, String author, boolean isReady){
        super();
        this.bookId = bookId;
        this.name = name;
        this.author = author;
        this.isReady = isReady;
    }
}
