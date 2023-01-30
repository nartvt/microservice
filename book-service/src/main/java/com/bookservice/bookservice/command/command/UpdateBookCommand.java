package com.bookservice.bookservice.command.command;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.axonframework.modelling.command.TargetAggregateIdentifier;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class UpdateBookCommand {
    @TargetAggregateIdentifier
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;

}
