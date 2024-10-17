# 實現一個 Map:

1. 面向高併發
2. 只存在插入和查詢操作 O(1)
3. 查詢時，如果 Key 存在，返回對應的 Value; 如果不存在，阻塞直到 k-v Pair 被放入，獲取 Value
4. 寫出真實代碼，不能有死鎖或 Panic