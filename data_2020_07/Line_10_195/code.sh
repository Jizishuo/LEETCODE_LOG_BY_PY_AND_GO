grep -n "" file.txt |grep -w '10' | cut -d: -f2

sed -n '10p' file.txt

awk '{if(NR==10){PRINT $0}}' file.txt


row_num=$(cat file.txt|wc -l)
echo $row_num
if [ $row_num -lt 10 ];then
  echo "no Enough 10 line"
else
  awk '{if{NR==10}{print $0}}' file.txt
fi