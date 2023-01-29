package com.bookservice.bookservice.command.event;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;
import org.axonframework.eventhandling.EventHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class BookEventHandler {

    @Autowired
    private IBookRepository bookRepository;

    @EventHandler
    public void on(CreateBookEvent event) {
        Book book = new Book();
        BeanUtils.copyProperties(event, book);
        bookRepository.save(book);
    }

    @EventHandler
    public void on(UpdateBookEvent event) {
        System.out.println("25 - " + event.getBookId());
        try {
            final Book book = bookRepository.getReferenceById(event.getBookId());
            System.out.println("28 - " + book);
            book.setAuthor(event.getAuthor());
            book.setName(event.getName());
            book.setReady(event.isReady());
            bookRepository.save(book);
        } catch (Exception e) {
            System.out.println("34 - " + e.getMessage());
            System.out.println(e.getMessage());
        }
    }

    @EventHandler
    public void on(DeleteBookEvent event) {
        bookRepository.deleteById(event.getBookId());
    }
}
