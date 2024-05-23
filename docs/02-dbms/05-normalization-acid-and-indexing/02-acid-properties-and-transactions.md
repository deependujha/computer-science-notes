# ACID properties & Transactions

## Transactions

- A unit of work done against the DB in a logical sequence.
- Sequence is very important in transaction.
- It is a **logical unit of work that contains one or more SQL statements**.
- The result of all these statements in a transaction either gets completed successfully (all the changes made to the database are permanent) or if at any point any failure happens it gets roll-backed (all the changes being done are undone).

!!! example
    For example, let's suppose we're sending some amount X from user A to user B. The sequence of steps are:

    - Read balance from A.
    - If balance is greater than or equal to X, subtract X from the balance and store it in buffer.
    - Write the updated balance in DB for user A.
    - Read balance from B.
    - Add X to B's balance in buffer.
    - Write the buffer in DB for user B.

    The above 6 sequences altogether comprise a single transaction.

---

## ACID Properties (`Atomic`, `Consistent`, `Isolation`, `Durable`)

![acid properties](../../images/dbms/ACID-Properties.jpg)

### Atomicity

!!! info ""
    - Either all operations of transaction are reflected properly in the DB, or none are.

---

### Consistency

!!! info ""
    - Integrity constraints must be maintained before and after transaction. (example, `balance can't be negative`)
    - DB must be consistent after transaction happens.

---

### Isolation

!!! info ""
    - Even though multiple transactions may execute concurrently, the system guarantees that, for every pair of transactions Ti and Tj, it appears to Ti that either Tj finished execution before Ti started, or Tj started execution after Ti finished. Thus, each transaction is unaware of other transactions executing concurrently in the system.
    - **Multiple transactions can happen in the system in isolation, without interfering each other**.
    - Multiple transactions can occur simultaneously, if they aren't interfering with each other, else, they will be executed one after other.

???+ example "Explanation on Isolation"
    - Let's say A has $1000 in his account and he transfers $50 to B. At the same time, from another device, A transfer B $50.
    - If Both transactions occur simultaneously, both transactions will read $1000 to be A's balance, and the updated balance in both case will be $950.
    - To tackle this, **isolation** is used.
    - Until one transaction completes, other won't even start.
    - This prevents unintended reads & writes.
    - It will use some sort of lock mechanism.

---

### Durability

!!! info ""
    - After transaction completes successfully, the changes it has made to the database persists, even if there are system failures.

???+ example "Explanation on Durability"
    - Suppose we have read balance from user A, subtracted X from A's balance, write the transaction, then, read B's balance, and added X to it. And, finally, write the updated balanced into the buffer. And, it will indicate a **Successful transaction**.
    - But, most of the time, we don't directly write to the DB. We store the updated write to be performed in the memory buffer. We do this because, it is typically a time-consuming process to go from execution mode to I/O mode. Hence, we typically store in the buffer, and from time to time, we flush it and commit to the disk (DB).
    - But, what if, when we responded the user with success, but we were yet to write to disk and the system crashed?
    - When we will restart system, memory would have been freed and flushed, so we lost the transaction, and hence the durability has been lost.
    - To achieve durability, DBMS have **recovery management component**, whose job is to write the successful transactions to the DB which have previously been marked to be successful.
    - For this, they generally emit logs, and when the system restarts, it tries to match with the log and check if the DB is at latest update, else, it writes missing transactions. Hence, Durability is achieved.

---

## Transaction states

![transaction states](../../images/dbms/transaction-states.png)

- Active state

!!! quote ""
    - The very first state of the life cycle of the transaction, all the read and write operations are being performed. If they execute without any error the T comes to Partially committed state. Although if any error occurs then it leads to a Failed state.

---

- Partially committed state

!!! quote ""
    - After transaction is executed the changes are saved in the buffer in the main memory. If the changes made are permanent on the DB then the state will transfer to the committed state and if there is any failure, the T will go to Failed state.

---

- Committed state

!!! quote ""
    - When updates are made permanent on the DB. Then the T is said to be in the committed state. Rollback can’t be done from the committed states. New consistent state is achieved at this stage.

---

- Failed state

!!! quote ""
    - When T is being executed and some failure occurs. Due to this it is impossible to continue the execution of the T.

---

- Aborted state

!!! quote ""
    - When T reaches the failed state, all the changes made in the buffer are reversed. After that the T rollback completely. T reaches abort state after rollback. DB’s state prior to the T is achieved.

---

- Terminated state

!!! quote ""
    - A transaction is said to have terminated if has either committed or aborted.
