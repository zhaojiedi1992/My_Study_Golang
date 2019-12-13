sync.Map
====================

map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。

sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，
因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。


.. literalinclude:: ../src/syncmap.go
   :encoding: utf-8
   :language: go
