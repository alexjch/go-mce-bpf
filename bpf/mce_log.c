#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24);
} mce_events SEC(".maps");

SEC("tracepoint/mce/mce_record")
int handle_mce(struct mce *args) {
    int pid = bpf_get_current_pid_tgid() >> 32;
    const char fmt_str[] = "Hello, world, from BPF! My PID is %d\n";
    bpf_trace_printk(fmt_str, sizeof(fmt_str), pid);
    bpf_ringbuf_output(&mce_events, args, sizeof(struct mce), 0);
    return 0;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";