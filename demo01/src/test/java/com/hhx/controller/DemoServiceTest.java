package com.hhx.controller;

import com.hhx.service.DemoService;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;

public class DemoServiceTest {

    private Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    private DemoService demoService;

    @Test
    public void task01() {
        long now = System.currentTimeMillis();
        logger.info("[task01][开始执行]");

        demoService.execute01();
        demoService.execute02();

        logger.info("[task01][结束执行，消耗时长 {} 毫秒]", System.currentTimeMillis() - now);
    }
}
