package com.hhx.controller;

import com.hhx.service.IMockTestService;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureMockMvc
public class MockTestControllerTest {

    @MockBean
    private IMockTestService mockTestService;

    @Test
    public void testGet() throws Exception{
        Mockito.when(mockTestService.get()).thenReturn("mock");
    }
}
