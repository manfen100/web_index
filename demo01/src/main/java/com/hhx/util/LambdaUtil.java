package com.hhx.util;

import java.util.ArrayList;
import java.util.List;

public class LambdaUtil {

    volatile List<Integer> integers = new ArrayList();
    // 基准测试1
    public int forEachLoopMaxInteger() {
        int max = Integer.MIN_VALUE;
        for (Object n : integers) {
            max = Integer.max(max, (Integer) n);
        }
        return max;
    }

    // 基准测试2
    public int lambdaMaxInteger() {
        return integers.stream().reduce(Integer.MIN_VALUE, (a, b) -> Integer.max(a, b));
    }
}
