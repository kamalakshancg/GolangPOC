package com.poc.controller;

import com.poc.entity.Order;
import com.poc.services.OrderService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class OrderController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Autowired
    private OrderService orderService;

    @GetMapping("/api/test1")
    public String test1() {
        return jdbcTemplate.queryForObject("SELECT 'pong'", String.class);
    }

    @GetMapping("/api/test2")
    public List<Order> test2() {
        return orderService.getOrderDetails();
    }

}
