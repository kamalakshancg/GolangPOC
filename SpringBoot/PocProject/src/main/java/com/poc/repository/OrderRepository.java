package com.poc.repository;

import com.poc.entity.Order;
import org.springframework.data.jpa.repository.JpaRepository;

import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
@Repository
public interface OrderRepository extends JpaRepository<Order, Integer> {

    @Query(value = "SELECT id, user_id, amount, status, description FROM orders ORDER BY id LIMIT 1000", nativeQuery = true)
    List<Order> getOrders();
}