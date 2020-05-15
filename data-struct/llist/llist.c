#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef void node_proc_fun_t(void *);
typedef int node_comp_fun_t(const void*, const void*);
typedef void LLIST_T;

LLIST_T *llist_new(int elmsize);
int llist_delete(LLIST_T *ptr);
LLIST_T *show(LLIST_T *ptr);
LLIST_T *showDebug(LLIST_T *ptr);

int llist_node_append(LLIST_T *ptr, const void *datap);
int llist_node_prepend(LLIST_T *ptr, const void *datap);

int llilst_travel(LLIST_T *ptr, node_proc_fun_t *proc);

void llist_node_delete(LLIST_T *ptr, node_comp_fun_t *comp, const void *key);
void* llist_node_find(LLIST_T *ptr, node_comp_fun_t* comp, const void* key);

struct node_st {
    void *datap;
    struct node_st *next, *prev;
};

struct llist_st {
    struct node_st head;
    int elmsize;
    int elmnr;
};

typedef int* dataType;

int main(){
    printf("hello\n");
    struct llist_st *ptr = llist_new(100);
    dataType data = malloc(100);
    *data = 10;
    llist_node_append(ptr, data);
    data = malloc(sizeof(dataType));
        *data = 10;
    llist_node_prepend(ptr, data);
    data = malloc(sizeof(dataType));
        *data = 10;
    llist_node_append(ptr, data);
    data = malloc(sizeof(dataType));
        *data = 10;
    llist_node_prepend(ptr, data);
    data = malloc(sizeof(dataType));
        *data = 10;
    llist_node_append(ptr, data);
    data = malloc(sizeof(dataType));
        *data = 10;
    llist_node_prepend(ptr, data);
    show(ptr);
    showDebug(ptr);
    return 0;
}

LLIST_T *llist_new(int elmsize){
    struct llist_st *newlist;
    newlist = malloc(sizeof(struct llist_st));
    if (newlist == NULL)
        return NULL;
    newlist->head.datap = NULL;
    newlist->head.next = &newlist->head;
    newlist->head.prev = &newlist->head;
    newlist->elmsize = elmsize;
    return (void *)newlist;
}

int llist_delete(LLIST_T* ptr) {
    struct llist_st *me = ptr;
    struct node_st *curr,*save;
    for (curr = me->head.next; curr != &me->head; curr = save) {
        save = curr->next;
        free(curr->datap);
        free(curr);
    }
    free(me);
    return 0;
}

int llist_node_append(LLIST_T *ptr, const void *datap) {
    struct llist_st *me = ptr;
    struct node_st *newnodep;
    newnodep = malloc(sizeof(struct node_st));
    if (newnodep == NULL) {
        return -1;
    }
    newnodep->datap = malloc(me->elmsize);
    if (newnodep->datap ==NULL){
        free(newnodep);
        return -1;
    }
    memcpy(newnodep->datap,datap,me->elmsize);
    me->head.prev->next = newnodep;
    newnodep->prev = me->head.prev;
    me->head.prev = newnodep;
    newnodep->next = &me->head;
    return 0;
}

int llist_node_prepend(LLIST_T *ptr, const void *datap) {
    struct llist_st *me = ptr;
    struct node_st *newnodep;
    newnodep  = malloc(sizeof(struct node_st));
    if (newnodep == NULL) return -1;
    newnodep->datap = malloc(me->elmsize);
    if (newnodep->datap == NULL) {
        free(newnodep);
        return -1;
    }
    memcpy(newnodep->datap, datap, me->elmsize);
    me->head.next->prev = newnodep;
    newnodep->next = me->head.next;
    me->head.next = newnodep;
    newnodep->prev = &me->head;
    return 0;
}

int llist_travel(LLIST_T *ptr, node_proc_fun_t *proc) {
    struct llist_st *me = ptr;
    struct node_st *curr;
    for(curr = me->head.next; curr != &me->head; curr = curr->next)proc(curr->datap);
    return 0;
}



LLIST_T *show(LLIST_T *ptr) {
    struct llist_st *me = ptr;
    printf("size: %d, elmnr: %d\n", me->elmsize, me->elmnr);
    //printf("debug   %x   %x   %x \n", me->head.prev, me->head.datap, me->head.next);
    struct node_st *curr;
    //note:用一个其他的变量，指针会出问题：struct llist_st head = me->head;
    for (curr = me->head.next; curr != &me->head; curr = curr->next) {
        printf("%d --> ", *(dataType)curr->datap);
    }
    printf("\n");
}

LLIST_T *showDebug(LLIST_T *ptr) {
    struct llist_st *me = ptr;
    printf("size: %d, elmnr: %d\n", me->elmsize, me->elmnr);
    struct node_st *curr;
    for (curr = me->head.next; curr != &me->head; curr = curr->next) {
        //printf("debug %0x  ",curr);
        printf("%0x %d %0x --> ",curr->prev, *(dataType)curr->datap, curr->next);
    }
    printf("\n");
}

