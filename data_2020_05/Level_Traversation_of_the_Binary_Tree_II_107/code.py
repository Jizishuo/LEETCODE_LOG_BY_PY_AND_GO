class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:
    def levelOrderBottom(self, root):
        # queue = []                                                  # 结果列表
        # cur = [root]                                                # 接下来要循环的当前层节点，存的是节点
        # while cur:                                                  # 当前层存在结点时
        #     cur_layer_val = []                                      # 初始化当前层结果列表为空，存的是val
        #     next_layer_node = []                                    # 初始化下一层结点列表为空
        #     for node in cur:                                        # 遍历当前层的每一个结点
        #         if node:                                            # 如果该结点不为空，则进行记录
        #             cur_layer_val.append(node.val)                  # 将该结点的值加入当前层结果列表的末尾
        #             next_layer_node.extend([node.left, node.right]) # 将该结点的左右孩子结点加入到下一层结点列表
        #     if cur_layer_val:                                       # 只要当前层结果列表不为空
        #         queue.insert(0, cur_layer_val)                      # 则把当前层结果列表插入到队列首端
        #     cur = next_layer_node                                   # 下一层的结点变成当前层，接着循环
        # return queue                                                # 返回结果队列

        queue = []
        vur = [root]
        while vur:
            val_list = []
            next_val_list = []
            for node in vur:
                if node:
                    val_list.append(node.val)
                    next_val_list.append([node.left, node.right])
                if val_list:
                    queue.insert(0, val_list)
                vur = next_val_list
            return queue

