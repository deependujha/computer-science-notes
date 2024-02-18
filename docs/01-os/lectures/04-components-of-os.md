# Components of Operating System

![linux architecture](../../images/os/architecture-of-linux.png)

## Components of OS 😈

1. **Kernel:** A kernel is that part of the operating system which interacts directly with
the hardware and performs the most crucial tasks.
    - Heart of OS/Core component
    - Very first part of OS to load on start-up.

2. **User space:** Where application software runs, apps don’t have privileged access to the
underlying hardware. It interacts with kernel.
    - GUI
    - CLI

!!! success "shell/ Command Interpreter"
    A **shell**, also known as a **command interpreter**, is that part of the operating system that **`receives commands from the users and gets them executed.`**

---

## Functions of Kernel 👮🏻‍♂️

!!! note "1. Process management"
    - Scheduling processes and threads on the CPUs.
    - Creating & deleting both user and system process.
    - Suspending and resuming processes
    - Providing mechanisms for process synchronization or process communication.

!!! note "2. Memory management"
    - Allocating and deallocating memory space as per need.
    - Keeping track of which part of memory are currently being used and by which process.

!!! note "3. File management"
    - Creating and deleting files.
    - Creating and deleting directories to organize files.
    - Mapping files into secondary storage.
    - Backup support onto a stable storage media.

!!! note "4. I/O management"
    - to manage and control I/O operations and I/O devices
    - Buffering (data copy between two devices), caching and spooling.
        - Within differing speed two jobs.
        - Eg. Print spooling and mail spooling.
    - Buffering
        - Within one job.
        - Eg. Youtube video buffering
    - Caching
        - Memory caching, Web caching etc.

??? failure "what is spooling?"
    - **Spooling** is a process in which data is temporarily held to be used and executed by a device, program or the system. It is a place where data is stored to be processed or printed later.
    - In modern systems, spooling is often used when data needs to be transferred from a program to a peripheral (such as a printer).

---

## Types of Kernels 🤖

### **1. Monolithic Kernel**
- All functions are in kernel itself.
- Bulky in size.
- Memory required to run is high.
- Less reliable, one module crashes -> whole kernel is down.
- High performance as communication is fast. (Less user mode, kernel mode overheads)
- Eg. **Linux, Unix, MS-DOS**.

---

### **2. Micro Kernel**
- Only major functions are in kernel.
    - Memory management.
    - Process management.
- File management and IO management are in User-space.
- smaller in size.
- More Reliable
- More stable
- Performance is slow.
- Overhead switching b/w user mode and kernel mode.
- Eg. **L4 Linux, Symbian OS, MINIX** etc.

---

### **3. Hybrid Kernel**

- Advantages of both worlds. (File management. in User space and rest in Kernel space. )
- Combined approach.
- Speed and design of mono.
- Modularity and stability of micro.
- Eg. **MacOS, Windows NT/7/10**
- IPC also happens but lesser overheads

---

### **4. Nano/Exo kernels**
- Smallest in size.
- Only basic functions.
- Rest in user space.
- Eg. **Singularity, EROS**

??? info "Nano/Exo kernels"
    a type of operating system kernel that has a minimal and streamlined design, with a focus on minimalism and efficiency. The goal of a nano kernel is to provide only the essential functions required for the operation of a system while delegating other functions to user-space processes.

---

## Communication b/w User mode and Kernel mode 📡
???+ danger "Q. How will communication happen between user mode and kernel mode?"
    -   **Inter process communication (IPC)**.
        - Two processes executing independently, having independent memory space (Memory protection), But some may need to communicate to work.
        - **Done by shared memory and message passing**.
