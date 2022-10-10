package com.hhx.service.impl;

import com.hhx.service.OrderService;
import io.seata.spring.annotation.GlobalTransactional;
import org.springframework.stereotype.Service;

@Service
public class OrderServiceImpl implements OrderService {

    @Override
    @GlobalTransactional//声明全局事务
    public Integer createOrder(Long userId, Long productId, Integer price) throws Exception {
        return null;
    }
}
