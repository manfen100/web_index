package com.hhx.util;
import java.math.BigDecimal;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;


/**
 * 一致性Hash的一种算法 高效低碰撞率
 */

public class Murmurs {

    /**
     * murmur hash算法实现
     */
    public static Long hash(byte[] key) {

        ByteBuffer buf = ByteBuffer.wrap(key);
        int seed = 0x1234ABCD;

        ByteOrder byteOrder = buf.order();
        buf.order(ByteOrder.LITTLE_ENDIAN);

        long m = 0xc6a4a7935bd1e995L;
        int r = 47;

        long h = seed ^ (buf.remaining() * m);

        long k;
        while (buf.remaining() >= 8) {
            k = buf.getLong();

            k *= m;
            k ^= k >>> r;
            k *= m;

            h ^= k;
            h *= m;
        }

        if (buf.remaining() > 0) {
            ByteBuffer finish = ByteBuffer.allocate(8).order(
                    ByteOrder.LITTLE_ENDIAN);
            // for big-endian version, do this first:
            // finish.position(8-buf.remaining());
            finish.put(buf).rewind();
            h ^= finish.getLong();
            h *= m;
        }

        h ^= h >>> r;
        h *= m;
        h ^= h >>> r;

        buf.order(byteOrder);
        return h;
    }

    public static Long hash(String key) {
        return hash(key.getBytes());
    }



    /**
     * Long转换成无符号长整型（C中数据类型）
     */
    public static BigDecimal readUnsignedLong(long value) {
        if (value >= 0)
            return new BigDecimal(value);
        long lowValue = value & 0x7fffffffffffffffL;
        return BigDecimal.valueOf(lowValue).add(BigDecimal.valueOf(Long.MAX_VALUE)).add(BigDecimal.valueOf(1));
    }

    /** 将hash值和100 取余，然后加1，计算得到1-100之间的数字*/
    public static int getHashValue(String id){
        BigDecimal amt =  Murmurs.hashUnsigned(id);
        BigDecimal[] results = amt.divideAndRemainder(BigDecimal.valueOf(100));
        // 此处可以加1操作，最终得到 1 到 100 之间的数值
        return results[1].intValue()+1;
    }

    /**
     * 返回无符号murmur hash值
     */
    public static BigDecimal hashUnsigned(String key) {
        return readUnsignedLong(hash(key));
    }
    public static BigDecimal hashUnsigned(byte[] key) {
        return readUnsignedLong(hash(key));
    }

    public static void main(String[] args) {

        System.out.println(getHashValue("2E2DFA89-496A-47FD-9941-DF1FC4E6484A"));
    }
}
