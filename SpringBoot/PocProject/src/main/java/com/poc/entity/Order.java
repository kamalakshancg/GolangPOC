package com.poc.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import jakarta.persistence.Transient;
import lombok.Getter;

import java.util.HashSet;
import java.util.Set;


@Entity
@Table(name = "orders")
public class Order {
    @Id
    private Integer id;
    private Double amount;
    private String status;
    private String description;
    private Integer userId;

    @Transient
    private Set<Item> items = new HashSet<>();

    protected Order(){
    }

    public Order(Integer id, Double amount, String status, String description, Integer userId) {
        this.id = id;
        this.amount = amount;
        this.status = status;
        this.description = description;
        this.userId = userId;
    }

    public void addItem(final Item item){
        this.items.add(item);
    }

    public Integer getId() {
        return id;
    }

    public Double getAmount() {
        return amount;
    }

    public String getStatus() {
        return status;
    }

    public String getDescription() {
        return description;
    }

    public Integer getUserId() {
        return userId;
    }

    public Set<Item> getItems() {
        return items;
    }
}