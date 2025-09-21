INSERT INTO students
(name, age , grade )
VALUES('张三', 20, '三年级', );


SELECT name, age , grade
FROM pmu.students
WHERE age > 18;


UPDATE  pmu.students SET  grade =  '四年级' where  name = '张三'


DELETE FROM pmu.students
WHERE age < 15;


START TRANSACTION;
  ##查询A余额
  select    id , balance
  FROM pmu.accounts
  WHERE  id = 'A' and balance >100;
  ##查询B余额
    select    id , balance
    FROM pmu.accounts
    WHERE  id = 'B' ;
##修改B余额
 UPDATE  pmu.accounts SET  balance += 100  where  name = 'B'
##修改A余额
UPDATE  pmu.accounts SET  balance -= 100  where  name = 'A'


INSERT INTO transactions
(id , from_account_id  , to_account_id ,amount  )
VALUES(1, 'A', 'B', 100);


COMMIT;
