package com.compressibleflowcalculator.Timestamps_Projects.Java.Models;

import java.time.OffsetDateTime;
import java.util.UUID;

import org.hibernate.annotations.ColumnDefault;

import jakarta.annotation.Generated;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;

@Entity
public class Notification {
    
    @Id
    @ColumnDefault("gen_random_uuid()")
    @jakarta.persistence.GeneratedValue
    private UUID id;

    private String message; 


    @Column(columnDefinition = "TIMESTAMP WITH TIME ZONE")

    private OffsetDateTime timestamp;

    //getters and setters
    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public OffsetDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(OffsetDateTime timestamp) {
        this.timestamp = timestamp;
    }

    // default constructor
    public Notification() {
    }

    // constructor
    public Notification( String message, OffsetDateTime timestamp) {

        this.message = message;
        this.timestamp = timestamp;
    }

    
    

}
