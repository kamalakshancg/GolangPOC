package com.poc.entity;

import jakarta.persistence.*;
import com.fasterxml.jackson.annotation.JsonBackReference;

@Entity
@Table(name = "items")
public class ItemEntity {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    public Integer id;
    public String productName;
    public Integer quantity;
    public Double price;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "order_id")
    @JsonBackReference
    public OrderEntity order;
}