package com.bookservice.bookservice.command.event;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;
import org.axonframework.eventhandling.EventHandler;
import org.jetbrains.annotations.NotNull;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.persistence.EntityNotFoundException;

@Component
public class BookEventHandler {
    private IBookRepository bookRepository;

    public BookEventHandler(IBookRepository bookRepository){
        this.bookRepository = bookRepository;
    }

    @EventHandler
    public void on(CreateBookEvent event) {
        Book book = new Book();
        BeanUtils.copyProperties(event, book);
        bookRepository.save(book);
    }

    @EventHandler
    public void on(UpdateBookEvent event) {
        final Book book = bookRepository.getReferenceById(event.getBookId());
        book.setAuthor(event.getAuthor());
        book.setName(event.getName());
        book.setReady(event.isReady());
        bookRepository.save(book);
    }

    @EventHandler
    public void on(DeleteBookEvent event) {
        if(event == null || event.getBookId() == null){
            return;
        }
        try {
            final Book book = bookRepository.getReferenceById(event.getBookId());
            if (book == null){
                return;
            }
            bookRepository.deleteById(book.getBookId());
        }catch (EntityNotFoundException e){
            System.out.println(e.getMessage());
        }
    }
}
