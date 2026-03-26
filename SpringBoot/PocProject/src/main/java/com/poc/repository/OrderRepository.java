package com.poc.repository;

import com.poc.entity.Order;
import org.springframework.data.jpa.repository.JpaRepository;

import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface OrderRepository extends JpaRepository<Order, Integer> {

    @Query(value = "SELECT * FROM orders ORDER BY id LIMIT 1000", nativeQuery = true)
    List<Order> findWideOrders();
}