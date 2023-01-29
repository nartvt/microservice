package com.bookservice.bookservice.command.data;

import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "books")
@Data
public class Book {
    @Id
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;
}
