package com.poc.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import java.util.List;
import java.util.Set;

import com.poc.entity.OrderEntity;
import com.poc.entity.UserEntity;
import com.poc.repository.OrderRepository;
import com.poc.repository.UserRepository;

@RestController
public class PocController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Autowired
    private OrderRepository orderRepository;

    @Autowired
    private UserRepository userRepository;

    @GetMapping("/api/test1")
    public String test1() {
        return jdbcTemplate.queryForObject("SELECT 'pong'", String.class);
    }

    @GetMapping("/api/test2")
    public List<OrderEntity> test2() {
        return orderRepository.findBy(PageRequest.of(0, 1000));
    }

    @GetMapping("/api/test3")
    public Set<UserEntity> test3() {
        return userRepository.findUsersWithOrdersAndItems();
    }
}