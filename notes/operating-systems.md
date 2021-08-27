3 major areas: virtualization, concurrency, persistence

operating system is software that is in charge of making sure that the system operates correctly and efficiently in an easy to use manner – programs can share memory, interact with devices, etc. It's also a resource manager - control program access to disk, memory and CPU time. 

virtualization: taking a physical resource and transform it into a virtual form – example would be the processor, memory, or disk. OS provides API in form of system calls / standard library. 

how does the OS virtualize the CPU? by creating the illusion that there are many CPUs available even though there is only one or a small set. 

OS policies - how the OS decides which program gets CPU time 

OS mechanisms

physical memory = array of bytes

OS virtualizes this by giving each program it's own private virtual address space - OS maps this to the actual physical memory of the machine

address space randomization - tool the OS uses to defend against "stack smashing attacks"

what happens when you call "open" system call? The OS translates this into various calls to (via a device driver) to the hardware 

journaling – allows system to recover in case of crash. 

what is the difference between a system call and a procedure call? making a system call is a formal transition to transfor control into the OS and raised the hardware priv level. Procedure calls run in "user mode" so they are restricted in what they can do - eg can't do I/O, or send a packet on the network. 

playing around with ubuntu 

- help by default seems to list all the shell commands available – builtins. help help 
- help = shell builtins like cd. man = references to online manuals for other commands like systems calls, standard library functions, etc. 

what is the Linux kernel? it's just a program that is loaded into memory on boot by a boot loader – example: GRUB. This program provides an interface of system calls to interact with it. It acts as a gatekeeper by enforcing privileges. Kernel is modular – base image allows you to boot, then you can add more. 

/boot contains kernel. 

possible questions for class

goals for this module

- build a solid mental model for where the hardware stops, the OS begins and how kernel land and userland are enforced - where one stops and the other begins. 
- understand different Linux distros – linux vs the linux flavors 
    - similarly - Windows vs MacOS vs Linux, history and foundations and offshoots. 
- does a bootloader come with the hardware or is that software too? what is the core operating system vs operating system modules? 

# class notes 1 

hypervisors - one physical machine, want to run multiple OSs on it, need a way to do that. So a hypervisor is an OS that runs OSses. 

KVM, Xen - many others - just forks of linux 

Multics introduced multitasking – precursors ran programs in series. 

Tree based system - with symlinks – hierarchy – other ideas – maybe graph lots of links. 

Filesystem isn't tightly coupled to the OS – you can make a choice about it. 

what comes with an OS - what do we expect? resource management, file system, device driver. 

what's the kernel? anything you have to go through a system call for. 

initd / systemd – first process to run and becomes parent process of all userland processes. 

slab allocation – kernel at boot time pre-allocates task structs that it can use for internal management so it doesn't have to do it dynamically. 

difference between a process and a thread? process has it's own slice of virtual memory allocated by the OS. They can each be independently scheduled by the OS (eg could run in parallel on separate CPUs). Whereas threads came later – want the same running process to have multiple execution units that can be scheduled independently, but could share memory. So single program broken up into execution units over CPUs. 

# class 2 prework - reading about processes 

program = instructions + data, process = running program

to run program, OS loads it from disk into program memory - memory is allocated for the runtime stack, then it jumps to the main() routine. 

stack - used for local vars, function params, return addresses. 

heap - used for explicitly requested dynamically allocated data - malloc + free. 

OS maintains process list to processes and process states, consisting of proc structs which maintain important info used for context switching – register contents, start of heap memory, size of memory, stack, state, etc. 

fork() syscall starts a new child process. 
wait() syscall waits for pid to exit.
exec() syscall run program separate from the caller - overwrites code segment - everything re-initialized. Call to exec never returns. 

fork vs exec 

stdin file descriptior = 0
stdout file descritor = 1
stderr file descriptor = 2

linkers https://www.lurklurk.org/linkers/linkers.html

– this reading was hard to follow, so switching over to CSAPP Ch 7 on Linking 

linking = collecting code and data into a single file that OS can load into memory and execute. Can happen at compile time or load time or run time. Allows decomposition into components that can be compiled separately. 

Questions from prework 

- session leader? 
- why did prtstat and pmap stack not line up? 
- building mental model of the kernel - can we talk about the boot loader and how that fits into it - where does the hardware stop and the OS begin? 
- cool if we start a google doc for topics for the final class?
- are we going to actually look at some kernel code during the course?

BIOS => GRUB (for Linux) 

Unikernel, Microkernels, modules and module loading, eBPF

executable file formats:https://en.wikipedia.org/wiki/Comparison_of_executable_file_formats

ELF: Executable and Linkable Format https://en.wikipedia.org/wiki/Executable_and_Linkable_Format

format for executable, shared library, or object files. 

segments: relevant at runtime (eg could have data segment with values and unitialized space, and then code)

and sections: relevant at link time. 

ELF Headers - e_ident - how to parse it - magic number, bit size of objects, byte ordering, etc. 

Mach-O
PE 

linkers: 

linking is combining code and data into a single file that can be loaded into memory and executed (but it can happen either at compile time, load time, or even at run time). They allow for separate compilation 

compiler driver hides what it is doing - 
preprocessing: creating single ascii text, compiling: ascii to assembly, assembling: assembly to relocatable object file and finally linking: combines relocatable object files. 

what's the linker doing? 
- resolve references between relocatable object files (resolving symbols (functions, global vars, static vars) to their definitions)
- determines the layout of the executable by relocating code and data to a memory location 

what's in an object file? It's just blocks of bytes that contain program code, program data, and data structures to guide the linker and loader. Linker just connects these blocks and decides where they will live in memory. 

relocatable object files vs executable object files? compilers and assemblers generate the former and linkers generate the latter. 

.bss – block started by symbol – uninitialized data - Better Save Space! 

.symtab - symbol table - info about functions and global vars 

- externals - global symbols that are referenced by the current module but are defined elsewhere


mangling - the compiler encodes duplicated symbols names (eg when you overload a function) as distinct names so that the linker doesn't get confused. 

notes from class
- what happens during preprocessing? macros are evaluated (replacing #defines)

shared libraries – dylib, dll, shared object file

static linking - 
runtime loading - 
dynamic loading - 

## reading: limited direct execution

time sharing - how can programs from userland run on the CPU while having the OS still maintain control of the overall system? 

Problem 1: restricting access to hardware 

OS is able to restrict access to hardware by forcing the userland processes to make sys calls through the OS (using system call numbers and calling conventions) - trap call allows userland process to handoff to OS and OS to hand back off to userland. Kernel creates a trap table at boot time which basically is a lookup of what hardware instructions to run based on which trap call is made. 

Problem 2: switching between processes 

how does OS get control back? 
coorperative approach: wait for system calls, trust process to behave
non-coop approach: timer interrupt that gives the OS back control every so often 

## reading: scheduling intro

performance (speed of job completion) is often at odds with fairness (giving equal amounts of CPU time to each process)

FIFO - simple, easy to implement, but depending on workload can be inefficient (eg if long running job starts before a short job)

convoy effect: short-jobs queuing behind long-jobs 

SJF - Shortest Job First: - improves avg turnaround time 

STCF - Shortest Time-to-Completion First  - 

measure performance: turnaround time (elapsed time to complete job) and response time (how long till it's scheduled on CPU)

Round Robin scheduling - run job for a time slice and switch to next job in queue - continue till all finished. 

when you introduce I/O, best case is using the CPU the whole time so switching off processes that are blocked by IO. 

two approaches on optimizing for different things - round robin = fairness and good response time but bad average turnaround time, SJF and STCF give better turnaround but worse response time. 

## reading: multi-level feedback queue

MLFQ: multi-level feedback queue - attempts to optimize turnaround time (hard because we don't know how long jobs will take) while also minimizing response time. 

How does MLFQ work?
- has multiple queues of different priority levels. 
- jobs with higher priority run first. jobs with equal priority share via round robin. 
- job priority is updated by the OS based on observed behavior 

to avoid starvation (long running jobs not getting CPU time because interactive processes are hogging it) - periodically boost priority of all jobs in the system and let them filter down again. 

to avoid processes gaming the scheduler, keep track of how much CPU time slice a process is using – if it uses it's full timeslice it still gets deprioritized. 

Post reading: 
what are the trade-offs between cooperative and preemptive scheduling? in coop scheduling, userland processes can monopolize CPU time by not making any system calls and just doing CPU intensive work. Preepmtive scheduling allows the OS to optimize for performance and fairness across multiple processes. 

What are the tradeoffs between the various CPU scheduling policies?

tradeoff is going to be between performance (responsiveness and turnaround time) and fairness. In the real world, the CPU doesn't know how long a process will take, so SJF and STCF aren't applicable as pure policies, but using MLFQ the OS can get a nice balance of fairness across processes and performance by learning about how each process is going to be using it's CPU time and how long its taking.  


# Class scratchpad on process execution and scheduling: 

unix + BSD - Berkeley software distribution - shell, network stack, c compiler, etc. 

BSD vs GNU versions of grep 

Mac => BSD lineage vs Ubuntu GNU lineage 

Lions' commentary on UNIX 

This is a longish paper on the Evolution of Unix Architecture https://ieeexplore.ieee.org/document/8704965

https://cs3210.cc.gatech.edu/r/unix6.pdf 

warsus.github.io/lions- 

the completely fair scheduler - task that we decide to schedule is based on how "overdue" it is for scheduling and the amount of time it gets on the CPU corresponds to that as well. 

# Virtual memory readings

## Address Spaces 

address space - running programs view of memory in the system (what it's allowed to write to). 
OS maps virtual memory exposed to the process to the real physical memory on the system which the process has no visibility into. 

stack = automatic memory because it's managed for you by the compiler. to add to stack just declare an variable. 
heap = long lived memory explicitly managed by the program - use malloc and free (nb these are not system calls, they are library calls) under the hood the memory manager is using system calls brk and sbrk to ask the OS to adjust program break (location of the end of the heap). mmap is another related syscall used to create an anon memory region associated with swap space. 

## Address Translation 
address transaltion often referred to as dynamic relocation - just translating virtual address to physical address. this can happen even after the process has started running. 

to do address translation, OS sets the base register and the hardware can use this to offset the virtual mem address to the physical address. Combined with limit/bounds register, this provides hardware support for memory protection. MMU - memory management unit - part of the CPU that helps with address translation. 

OS tracks available memory so it can allocate memory to processes. Could be other data structures, but free list is one. 

Hardware support: provides base and bounds registers, and a "processor status word" to indicate kernel or user mode. Also provides area of memory to call OS exception handler when out of bounds exception runs. 

Qs 
- why would the OS move a processes memory? just to fit more stuff (eg de-fragment - could this happen though with fixed size pages?) A: yes, it's to avoid fragmentation. 

## segmentation
in order to avoid internal fragmentation, put each section (code, heap, stack) at different places in physical memory so that they don't write toward each other. Thus a segmentation fault is when you try and access memory outside of the segment your process is allowed to access. 

## paging 

segmentation - chopping things up into variable sized pieces but this leads to fragmentation 
paging - chopping things up into fixed-sized pieces, physical memory broken into page frames. 

this provides more flexibility and simplicity 

OS records how virtual memory maps to physical memory using a per-process data structure called a page table. 

A virtual address is broken into the virtual page number (2 most significant bits) and an offset. The VPN is translated into a physical page number in the page table, and the offset tells you the byte within the page that you want. 

So you need N significant bits depending on how many virtual pages you have per process. 

basic implementation on modern machines would take up a bunch of memory just for address translation and would also be slow, many steps to look it up in main memory. 

how to make this faster? TLB: translation-lookaside buffer, part of MMU and basically a hardware cache of popular virtual address lookups - aka address translation cache.

CISC vs RISC - Complex Instruction Set Computing vs Reduced Instruction Set Computing – simple hardware primitives or advanced – give control to the compiler so it can figure out how to make performant code or expose a complex API to it. 

cache replacement: least recently used LRU, or random policy. 

question: how does the memory in the TLB and in the page table get marked as unallocated once process terminates? I think this would just be a case of the OS itself not caring anymore about the process so even though the data is there it's not tracked and can be overwritten. 

how does the OS manage "fairness" in TLB hits? eg what if a process is switched onto the CPU and makes a ton of memory accesses, replacing cached values for the previous process. 