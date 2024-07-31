package com.compressibleflowcalculator.Timestamps_Projects.Java.Models;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.time.OffsetDateTime;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;

/**
 * 
 * This will be used with Entities that have a timestamp
 * 
 * JPA Model
 * 
 */


@Entity
 public class ModelWithTimestamp {
     
    //Should Be JSON Serialized

    @Id
    @JsonProperty("name")
    public String name;

    //Time Stamp, Postgres Definition is TIMESTAMP WITH TIME ZONE
    
    @JsonProperty("timestamp")
    // pattern be RFC 1132
    @JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd'T'HH:mm:ss.SSSZ")
    @Column(columnDefinition = "TIMESTAMP WITH TIME ZONE")
    @org.springframework.beans.factory.annotation.Value("${spring.jpa.properties.hibernate.jdbc.time_zone}")
    public OffsetDateTime  timestamp;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public OffsetDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(OffsetDateTime timestamp) {
        this.timestamp = timestamp;
    }
     

    public ModelWithTimestamp(String name, OffsetDateTime timestamp) {
        this.name = name;
        this.timestamp = timestamp;
    }

    public ModelWithTimestamp() {
    }

 }