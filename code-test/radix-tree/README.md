
该代码实现了一个基于Radix树的单词存储方式。例如

给定一组单词： "opening", "open", "book", "books", "blue"

依次执行insert 添加到一棵Radix树后：
```shell

                        ''
                    /        \
                'b'          'open'
               /   \         /    \
            'ook'   'lue'   ''    'ing'
            /    \    /           /
           ''    's' ''          ''
                 /
               ''
```
其中：
* 空的叶节点表示 路径节点本身也是一个词，比如 book，相反，b 没有空叶节点，表示b不是一个单词
* 树的维护遵循radix树的压缩特点，如果删除了open，则树的右半部分变为 '' -- 'opening' - ''，即 open 和 ing合并。
* 路径节点中，只有根节点可能是 '' 空节点。
