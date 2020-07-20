class TripleInOne:

    def __init__(self, stackSize: int):
        self.tripstack = [[] for i in range(3)]
        self.stacksize = stackSize

    def push(self, stackNum: int, value: int) -> None:
        if len(self.tripstack[stackNum]) < self.stacksize:
            self.tripstack[stackNum].append(value)

    def pop(self, stackNum: int) -> int:
        if self.tripstack[stackNum] != []:
            cache = self.tripstack[stackNum][-1]
            self.tripstack[stackNum] = self.tripstack[stackNum][:-1]
            return cache
        else:return -1

    def peek(self, stackNum: int) -> int:
        if self.tripstack[stackNum] !=[]:
            cache = self.tripstack[stackNum][-1]
            return cache
        else:return -1

    def isEmpty(self, stackNum: int) -> bool:
        return self.tripstack[stackNum]==[]


# Your TripleInOne object will be instantiated and called as such:
obj = TripleInOne(stackSize)
obj.push(stackNum,value)
param_2 = obj.pop(stackNum)
param_3 = obj.peek(stackNum)
param_4 = obj.isEmpty(stackNum)