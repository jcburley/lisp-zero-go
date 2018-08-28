/*
	Package main - transpiled by c2go version: v0.25.2 Dubnium 2018-06-29

	If you have found any issues, please raise an issue at:
	https://github.com/elliotchance/c2go/
*/

// Warning (RecordDecl):  /usr/include/libio.h:144 : could not lookup type definition for : _IO_FILE
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/waitstatus.h:75 : Error : name of FieldDecl is empty
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/waitstatus.h:89 : Error : name of FieldDecl is empty
// Warning (TransparentUnionAttr):  /usr/include/stdlib.h:71 : could not parse &{32166928 {/usr/include/stdlib.h 71 0 35 0 } []}
// Warning (FieldDecl):  /usr/include/stdlib.h:69 : Avoid struct `union wait *` in FieldDecl
// Warning (RecordDecl):  /usr/include/stdlib.h:67 : could not determine the size of type `union __WAIT_STATUS` for that reason: Cannot determine sizeof : |union __WAIT_STATUS|. err = Cannot determine sizeof : |union wait *|. err = error in union
// Error (RecordDecl):  /usr/include/stdlib.h:67 : Cannot determine sizeof : |union __WAIT_STATUS|. err = Cannot determine sizeof : |union wait *|. err = error in union
// Warning (EnumDecl):  /usr/include/ctype.h:46 : Add support of continues counter for type : *ast.ParenExpr
// Warning: using unsafe slice cast to convert from []byte to []byte

package main

import "os"
import "github.com/elliotchance/c2go/linux"
import "github.com/elliotchance/c2go/noarch"
import "unsafe"

type size_t uint32
type __u_char uint8
type __u_short uint16
type __u_int uint32
type __u_long uint32
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int32
type __uint32_t uint32
type __int64_t int32
type __uint64_t uint32
type __quad_t int32
type __u_quad_t uint32
type __dev_t uint32
type __uid_t uint32
type __gid_t uint32
type __ino_t uint32
type __ino64_t uint32
type __mode_t uint32
type __nlink_t uint32
type __off_t int32
type __off64_t int32
type __pid_t int32
type __fsid_t struct {
	__val [2]int32
}
type __clock_t int32
type __rlim_t uint32
type __rlim64_t uint32
type __id_t uint32
type __time_t int32
type __useconds_t uint32
type __suseconds_t int32
type __daddr_t int32
type __key_t int32
type __clockid_t int32
type __timer_t unsafe.Pointer
type __blksize_t int32
type __blkcnt_t int32
type __blkcnt64_t int32
type __fsblkcnt_t uint32
type __fsblkcnt64_t uint32
type __fsfilcnt_t uint32
type __fsfilcnt64_t uint32
type __fsword_t int32
type __ssize_t int32
type __syscall_slong_t int32
type __syscall_ulong_t uint32
type __loff_t __off64_t
type __qaddr_t *__quad_t
type __caddr_t *byte
type __intptr_t int32
type __socklen_t uint32
type FILE _IO_FILE
type __FILE _IO_FILE
type BSunionSatSSusrSincludeSwcharPhD85D3E struct{ memory unsafe.Pointer }

func (unionVar *BSunionSatSSusrSincludeSwcharPhD85D3E) copy() BSunionSatSSusrSincludeSwcharPhD85D3E {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion BSunionSatSSusrSincludeSwcharPhD85D3E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *BSunionSatSSusrSincludeSwcharPhD85D3E) __wch() *uint32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint32)(unionVar.memory)
}
func (unionVar *BSunionSatSSusrSincludeSwcharPhD85D3E) __wchb() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}

type __mbstate_t struct {
	__count int32
	__value BSunionSatSSusrSincludeSwcharPhD85D3E
}
type _G_fpos_t struct {
	__pos   __off_t
	__state int64
}
type _G_fpos64_t struct {
	__pos   __off64_t
	__state int64
}
type va_list int64
type __gnuc_va_list int64
type _IO_jump_t struct {
}
type _IO_lock_t interface{}
type _IO_marker struct {
	_next *_IO_marker
	_sbuf *_IO_FILE
	_pos  int32
}
type __codecvt_result int32

const ( // Warning (RecordDecl):  /usr/include/libio.h:144 : could not lookup type definition for : _IO_FILE
	__codecvt_ok      __codecvt_result = 0
	__codecvt_partial                  = 1
	__codecvt_error                    = 2
	__codecvt_noconv                   = 3
)

type _IO_FILE struct {
	_flags          int32
	_IO_read_ptr    *byte
	_IO_read_end    *byte
	_IO_read_base   *byte
	_IO_write_base  *byte
	_IO_write_ptr   *byte
	_IO_write_end   *byte
	_IO_buf_base    *byte
	_IO_buf_end     *byte
	_IO_save_base   *byte
	_IO_backup_base *byte
	_IO_save_end    *byte
	_markers        *_IO_marker
	_chain          *_IO_FILE
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  int8
	_shortbuf       [1]byte
	_lock           *_IO_lock_t
	_offset         __off64_t
	__pad1          unsafe.Pointer
	__pad2          unsafe.Pointer
	__pad3          unsafe.Pointer
	__pad4          unsafe.Pointer
	__pad5          size_t
	_mode           int32
	_unused2        [20]byte
}
type _IO_FILE_plus struct {
}
type __io_read_fn func(unsafe.Pointer, *byte, size_t) __ssize_t
type __io_write_fn func(unsafe.Pointer, *byte, size_t) __ssize_t
type __io_seek_fn func(unsafe.Pointer, *__off64_t, int32) int32
type __io_close_fn func(unsafe.Pointer) int32
type cookie_read_function_t __io_read_fn
type cookie_write_function_t __io_write_fn
type cookie_seek_function_t __io_seek_fn
type cookie_close_function_t __io_close_fn
type _IO_cookie_io_functions_t struct {
	read  __io_read_fn
	write __io_write_fn
	seek  __io_seek_fn
	close __io_close_fn
}
type cookie_io_functions_t _IO_cookie_io_functions_t
type _IO_cookie_file struct {
}
type off_t __off_t
type off64_t __off64_t
type ssize_t __ssize_t
type fpos_t _G_fpos_t
type fpos64_t _G_fpos64_t

var stdin *noarch.File

var stdout *noarch.File

var stderr *noarch.File

type obstack struct {
}
type wchar_t int32

const P_ALL int32 = 0
const P_PID int32 = 1
const P_PGID int32 = 2

type idtype_t int32
type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD69D5E struct {
	__w_termsig  uint32
	__w_coredump uint32
	__w_retcode  uint32
}
type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD84D5E struct {
	__w_stopval uint32
	__w_stopsig uint32
}
type wait struct{ memory unsafe.Pointer }

func (unionVar *wait) copy() wait {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion wait
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *wait) w_status() *int32 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *wait) __wait_terminated() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD69D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD69D5E)(unionVar.memory)
}
func (unionVar *wait) __wait_stopped() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD84D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSwaitstatusPhD84D5E)(unionVar.memory)
}

type div_t struct {
	quot int32
	rem  int32
}
type ldiv_t struct {
	quot int32
	rem  int32
}
type lldiv_t struct {
	quot int64
	rem  int64
}
type __locale_t int32
type locale_t __locale_t
type u_char __u_char
type u_short __u_short
type u_int __u_int
type u_long __u_long
type quad_t __quad_t
type u_quad_t __u_quad_t
type fsid_t __fsid_t
type loff_t __loff_t
type ino_t __ino_t
type ino64_t __ino64_t
type dev_t __dev_t
type gid_t __gid_t
type mode_t __mode_t
type nlink_t __nlink_t
type uid_t __uid_t
type pid_t __pid_t
type id_t __id_t
type daddr_t __daddr_t
type caddr_t __caddr_t
type key_t __key_t
type clock_t __clock_t
type time_t __time_t
type clockid_t __clockid_t
type timer_t __timer_t
type useconds_t __useconds_t
type suseconds_t __suseconds_t
type ulong uint32
type ushort uint16
type uint uint32
type int8_t int8
type int16_t int16
type int32_t int32
type int64_t int32
type u_int8_t uint8
type u_int16_t uint16
type u_int32_t uint32
type u_int64_t uint32
type register_t int32
type __sig_atomic_t int32
type __sigset_t struct {
	__val [16]uint32
}
type sigset_t __sigset_t
type timespec struct {
	tv_sec  __time_t
	tv_nsec __syscall_slong_t
}
type timeval struct {
	tv_sec  __time_t
	tv_usec __suseconds_t
}
type __fd_mask int32
type fd_set struct {
	fds_bits [16]__fd_mask
}
type fd_mask __fd_mask
type blksize_t __blksize_t
type blkcnt_t __blkcnt_t
type fsblkcnt_t __fsblkcnt_t
type fsfilcnt_t __fsfilcnt_t
type blkcnt64_t __blkcnt64_t
type fsblkcnt64_t __fsblkcnt64_t
type fsfilcnt64_t __fsfilcnt64_t
type pthread_t uint32
type pthread_attr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_attr_t) copy() pthread_attr_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_attr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_attr_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_attr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type __pthread_internal_list struct {
	__prev *__pthread_internal_list
	__next *__pthread_internal_list
}
type __pthread_list_t __pthread_internal_list
type __pthread_mutex_s struct {
	__lock    int32
	__count   uint32
	__owner   int32
	__nusers  uint32
	__kind    int32
	__spins   int16
	__elision int16
	__list    __pthread_list_t
}
type pthread_mutex_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutex_t) copy() pthread_mutex_t {
	var buffer [40]byte
	for i := range buffer {
		buffer[i] = (*((*[40]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutex_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutex_t) __data() *__pthread_mutex_s {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__pthread_mutex_s)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __size() *[40]byte {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[40]byte)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_mutexattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutexattr_t) copy() pthread_mutexattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutexattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutexattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_mutexattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD141D3E struct {
	__lock          int32
	__futex         uint32
	__total_seq     uint64
	__wakeup_seq    uint64
	__woken_seq     uint64
	__mutex         unsafe.Pointer
	__nwaiters      uint32
	__broadcast_seq uint32
}
type pthread_cond_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_cond_t) copy() pthread_cond_t {
	var buffer [72]byte
	for i := range buffer {
		buffer[i] = (*((*[72]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_cond_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_cond_t) __data() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD141D3E {
	if unionVar.memory == nil {
		var buffer [72]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD141D3E)(unionVar.memory)
}
func (unionVar *pthread_cond_t) __size() *[48]byte {
	if unionVar.memory == nil {
		var buffer [72]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[48]byte)(unionVar.memory)
}
func (unionVar *pthread_cond_t) __align() *int64 {
	if unionVar.memory == nil {
		var buffer [72]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int64)(unionVar.memory)
}

type pthread_condattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_condattr_t) copy() pthread_condattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_condattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_condattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_condattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_key_t uint32
type pthread_once_t int32
type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD177D3E struct {
	__lock              int32
	__nr_readers        uint32
	__readers_wakeup    uint32
	__writer_wakeup     uint32
	__nr_readers_queued uint32
	__nr_writers_queued uint32
	__writer            int32
	__shared            int32
	__rwelision         int8
	__pad1              [7]uint8
	__pad2              uint32
	__flags             uint32
}
type pthread_rwlock_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlock_t) copy() pthread_rwlock_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlock_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlock_t) __data() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD177D3E {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSpthreadtypesPhD177D3E)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_rwlockattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlockattr_t) copy() pthread_rwlockattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlockattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlockattr_t) __size() *[8]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[8]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlockattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_spinlock_t int32
type pthread_barrier_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrier_t) copy() pthread_barrier_t {
	var buffer [32]byte
	for i := range buffer {
		buffer[i] = (*((*[32]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrier_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrier_t) __size() *[32]byte {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[32]byte)(unionVar.memory)
}
func (unionVar *pthread_barrier_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_barrierattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrierattr_t) copy() pthread_barrierattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrierattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrierattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_barrierattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type random_data struct {
	fptr      *int32_t
	rptr      *int32_t
	state     *int32_t
	rand_type int32
	rand_deg  int32
	rand_sep  int32
	end_ptr   *int32_t
}
type drand48_data struct {
	__x     [3]uint16
	__old_x [3]uint16
	__c     uint16
	__init  uint16
	__a     uint64
}
type __compar_fn_t func(unsafe.Pointer, unsafe.Pointer) int32
type comparison_fn_t __compar_fn_t
type __compar_d_fn_t func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32

const // Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/waitstatus.h:75 : Error : name of FieldDecl is empty
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/waitstatus.h:89 : Error : name of FieldDecl is empty
// Warning (TransparentUnionAttr):  /usr/include/stdlib.h:71 : could not parse &{32166928 {/usr/include/stdlib.h 71 0 35 0 } []}
// Warning (FieldDecl):  /usr/include/stdlib.h:69 : Avoid struct `union wait *` in FieldDecl
// Warning (RecordDecl):  /usr/include/stdlib.h:67 : could not determine the size of type `union __WAIT_STATUS` for that reason: Cannot determine sizeof : |union __WAIT_STATUS|. err = Cannot determine sizeof : |union wait *|. err = error in union
// Error (RecordDecl):  /usr/include/stdlib.h:67 : Cannot determine sizeof : |union __WAIT_STATUS|. err = Cannot determine sizeof : |union wait *|. err = error in union
_ISupper uint16 = ((1 << 0) << 8)
const // Warning (EnumDecl):  /usr/include/ctype.h:46 : Add support of continues counter for type : *ast.ParenExpr
_ISlower uint16 = ((1 << 1) << 8)
const _ISalpha uint16 = ((1 << 2) << 8)
const _ISdigit uint16 = ((1 << 3) << 8)
const _ISxdigit uint16 = ((1 << 4) << 8)
const _ISspace uint16 = ((1 << 5) << 8)
const _ISprint uint16 = ((1 << 6) << 8)
const _ISgraph uint16 = ((1 << 7) << 8)
const _ISblank uint16 = ((1 << 8) >> 8)
const _IScntrl uint16 = ((1 << 9) >> 8)
const _ISpunct uint16 = ((1 << 10) >> 8)
const _ISalnum uint16 = ((1 << 11) >> 8)

type map_node_t map_node_s
type map_base_s struct {
	buckets  **map_node_t
	nbuckets uint32
	nnodes   uint32
}
type map_base_t map_base_s
type map_iter_s struct {
	bucketidx uint32
	node      *map_node_t
}
type map_iter_t map_iter_s
type map_void_t struct {
	base map_base_t
	ref  *unsafe.Pointer
	tmp  unsafe.Pointer
}
type map_str_t struct {
	base map_base_t
	ref  **byte
	tmp  *byte
}
type map_int_t struct {
	base map_base_t
	ref  *int32
	tmp  int32
}
type map_char_t struct {
	base map_base_t
	ref  *byte
	tmp  byte
}
type map_float_t struct {
	base map_base_t
	ref  *float32
	tmp  float32
}
type map_double_t struct {
	base map_base_t
	ref  *float64
	tmp  float64
}
type map_node_s struct {
	hash  uint32
	value unsafe.Pointer
	next  *map_node_t
}

// map_hash - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:122
/* A Lisp version 0 interpreter
   Copyright (c) 2010 James Craig Burley <james@jcb-sc.com>

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2 of the License , or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program. If not, write to the Free Software
   Foundation, 675 Mass Ave, Cambridge, MA 02139, USA.
*/ //
/* Cloned and wildly modified from
   http://www.sonoma.edu/users/l/luvisi/sl3.c by Andru Luvisi.

   Zero version implements "first" version of Lisp via Lisp code on
   top of builtin functions quote, atom, eq, cons, car, cdr, and cond,
   along with builtin recursive evaluations of the arguments to those
   functions where appropriate.

   This version incorporate the map.c source file, to comprise a
   single source file, easing use of c2go to convert it to Go and
   build it from there.
*/ //
/**
 * Copyright (c) 2014 rxi
 *
 * This library is free software; you can redistribute it and/or modify it
 * under the terms of the MIT license. See LICENSE for details.
 */ //
/* char key[]; */ //
/* char value[]; */ //
//
func map_hash(str *byte) uint32 {
	var hash uint32 = uint32(int32(5381))
	for *str != 0 {
		hash = hash<<uint64(int32(5)) + hash ^ uint32(*func() *byte {
			defer func() {
				str = ((*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + (uintptr)(1)*unsafe.Sizeof(*str))))
			}()
			return str
		}())
	}
	return hash
}

// map_newnode - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:131
func map_newnode(key *byte, value unsafe.Pointer, vsize int32) *map_node_t {
	var node *map_node_t
	var ksize int32 = noarch.Strlen(key) + int32(uint32(int32(1)))
	var voffset int32 = int32(uint32(ksize) + (8-uint32(ksize))%8)
	node = (*map_node_t)(noarch.Malloc(int32(24 + uint32(voffset) + uint32(vsize))))
	if node == nil {
		return nil
	}
	noarch.Memcpy(unsafe.Pointer(((*map_node_t)(unsafe.Pointer(uintptr(unsafe.Pointer(node)) + (uintptr)(int32(1))*unsafe.Sizeof(*node))))), unsafe.Pointer(key), int32(uint32(ksize)))
	(*node).hash = map_hash(key)
	(*node).value = unsafe.Pointer(((*byte)(func() unsafe.Pointer {
		tempVar := (*byte)(unsafe.Pointer(((*map_node_t)(unsafe.Pointer(uintptr(unsafe.Pointer(node)) + (uintptr)(int32(1))*unsafe.Sizeof(*node))))))
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(voffset)*unsafe.Sizeof(*tempVar))
	}())))
	noarch.Memcpy((*node).value, value, int32(uint32(vsize)))
	return node
}

// map_bucketidx - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:145
/* If the implementation is changed to allow a non-power-of-2 bucket count,
 * the line below should be changed to use mod instead of AND */ //
//
func map_bucketidx(m *map_base_t, hash uint32) int32 {
	return int32(hash & ((*m).nbuckets - uint32(int32(1))))
}

// map_addnode - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:152
func map_addnode(m *map_base_t, node *map_node_t) {
	var n int32 = map_bucketidx(m, (*node).hash)
	(*node).next = *((**map_node_t)(func() unsafe.Pointer {
		tempVar := (*m).buckets
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(n)*unsafe.Sizeof(*tempVar))
	}()))
	*((**map_node_t)(func() unsafe.Pointer {
		tempVar := (*m).buckets
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(n)*unsafe.Sizeof(*tempVar))
	}())) = node
}

// map_resize - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:159
/* Chain all nodes together */ //
/* Reset buckets */ //
/* Re-add nodes to buckets */ //
/* Return error code if realloc() failed */ //
//
func map_resize(m *map_base_t, nbuckets int32) int32 {
	var nodes *map_node_t
	var node *map_node_t
	var next *map_node_t
	var buckets **map_node_t
	var i int32
	nodes = nil
	i = int32((*m).nbuckets)
	for func() int32 {
		defer func() {
			i -= 1
		}()
		return i
	}() != 0 {
		node = *((**map_node_t)(func() unsafe.Pointer {
			tempVar := ((*m).buckets)
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
		}()))
		for node != nil {
			next = (*node).next
			(*node).next = nodes
			nodes = node
			node = next
		}
	}
	buckets = (**map_node_t)(noarch.Malloc(int32(8 * uint32(nbuckets))))
	if buckets != nil {
		(*m).buckets = buckets
		(*m).nbuckets = uint32(nbuckets)
	}
	if (*m).buckets != nil {
		noarch.Memset(unsafe.Pointer((*m).buckets), int32(0), int32(8*(*m).nbuckets))
		node = nodes
		for node != nil {
			next = (*node).next
			map_addnode(m, node)
			node = next
		}
	}
	return func() int32 {
		if (map[bool]int32{false: 0, true: 1}[buckets == nil]) != 0 {
			return -int32(1)
		} else {
			return int32(0)
		}
	}()
}

// map_getref - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:196
func map_getref(m *map_base_t, key *byte) **map_node_t {
	var hash uint32 = map_hash(key)
	var next **map_node_t
	if (*m).nbuckets > uint32(int32(0)) {
		next = &*((**map_node_t)(func() unsafe.Pointer {
			tempVar := (*m).buckets
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(map_bucketidx(m, hash))*unsafe.Sizeof(*tempVar))
		}()))
		for *next != nil {
			if (*(*next)).hash == hash && noarch.NotInt32(noarch.Strcmp((*byte)(unsafe.Pointer(((*map_node_t)(func() unsafe.Pointer {
				tempVar := *next
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
			}())))), key)) != 0 {
				return next
			}
			next = &(*(*next)).next
		}
	}
	return nil
}

// map_deinit_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:212
func map_deinit_(m *map_base_t) {
	var next *map_node_t
	var node *map_node_t
	var i int32
	i = int32((*m).nbuckets)
	for func() int32 {
		defer func() {
			i -= 1
		}()
		return i
	}() != 0 {
		node = *((**map_node_t)(func() unsafe.Pointer {
			tempVar := (*m).buckets
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(i)*unsafe.Sizeof(*tempVar))
		}()))
		for node != nil {
			next = (*node).next
			noarch.Free(unsafe.Pointer(node))
			node = next
		}
	}
	noarch.Free(unsafe.Pointer((*m).buckets))
}

// map_get_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:228
func map_get_(m *map_base_t, key *byte) unsafe.Pointer {
	var next **map_node_t = map_getref(m, key)
	return func() unsafe.Pointer {
		if next != nil {
			return (*(*next)).value
		} else {
			return (nil)
		}
	}()
}

// map_set_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:234
/* Find & replace existing node */ //
/* Add new node */ //
//
func map_set_(m *map_base_t, key *byte, value unsafe.Pointer, vsize int32) int32 {
	var n int32
	var err int32
	var next **map_node_t
	var node *map_node_t
	next = map_getref(m, key)
	if next != nil {
		noarch.Memcpy((*(*next)).value, value, int32(uint32(vsize)))
		return int32(0)
	}
	node = map_newnode(key, value, vsize)
	if node == nil {
		goto fail
	}
	if (*m).nnodes >= (*m).nbuckets {
		n = int32(func() uint32 {
			if (map[bool]int32{false: 0, true: 1}[(*m).nbuckets > uint32(int32(0))]) != 0 {
				return ((*m).nbuckets << uint64(int32(1)))
			} else {
				return uint32(int32(1))
			}
		}())
		err = map_resize(m, n)
		if err != 0 {
			goto fail
		}
	}
	map_addnode(m, node)
	(*m).nnodes += uint32(int32(1))
	return int32(0)
fail:
	;
	if node != nil {
		noarch.Free(unsafe.Pointer(node))
	}
	return -int32(1)
}

// map_remove_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:260
func map_remove_(m *map_base_t, key *byte) {
	var node *map_node_t
	var next **map_node_t = map_getref(m, key)
	if next != nil {
		node = *next
		*next = (*(*next)).next
		noarch.Free(unsafe.Pointer(node))
		(*m).nnodes -= uint32(int32(1))
	}
}

// map_iter_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:272
func map_iter_() map_iter_t {
	var iter map_iter_t
	iter.bucketidx = uint32(int32(2147483647))*uint32(2) + uint32(1)
	iter.node = nil
	return map_iter_t(iter)
}

// map_next_ - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:280
func map_next_(m *map_base_t, iter *map_iter_t) *byte {
	if (*iter).node == nil || (func() *map_node_t {
		(*iter).node = (*(*iter).node).next
		return (*iter).node
	}()) == nil {
		for {
			if func() uint32 {
				tempVar := &(*iter).bucketidx
				*tempVar += 1
				return *tempVar
			}() >= (*m).nbuckets {
				return nil
			}
			(*iter).node = *((**map_node_t)(func() unsafe.Pointer {
				tempVar := (*m).buckets
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32((*iter).bucketidx))*unsafe.Sizeof(*tempVar))
			}()))
			if noarch.NotInt32((map[bool]int32{false: 0, true: 1}[(*iter).node == nil])) != 0 {
				break
			}
		}
	}
	return (*byte)(unsafe.Pointer(((*map_node_t)(func() unsafe.Pointer {
		tempVar := (*iter).node
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
	}()))))
}

var quiet int8 = int8((int8(int32(0))))
var allocations int64 = int64(int32(0))
var allocations_total int64 = int64(int32(0))

// string_duplicate - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:333
/* Heap. */ //
//
func string_duplicate(str *byte) *byte {
	var l size_t = size_t((uint32(noarch.Strlen(str) + int32(uint32(int32(1))))))
	var dup *byte = (*byte)(noarch.Malloc(int32(uint32((size_t(l))))))
	if dup == nil {
		for {
			var // Warning: using unsafe slice cast to convert from []byte to []byte
			m *byte = (&[]byte("insufficient memory for duplicating a string\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(998)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	noarch.Memcpy(unsafe.Pointer(dup), unsafe.Pointer(str), int32(uint32((size_t(l)))))
	allocations += 1
	allocations_total += int64(uint64(noarch.Strlen(str)))
	return dup
}

type compiled_fn func(*byte, *Object_s, *Object_s) *Object_s
type Symbol_s struct {
	name *byte
}
type Cdr_u struct{ memory unsafe.Pointer }

func (unionVar *Cdr_u) copy() Cdr_u {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion Cdr_u
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *Cdr_u) obj() **Object_s {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (**Object_s)(unionVar.memory)
}
func (unionVar *Cdr_u) sym() **Symbol_s {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (**Symbol_s)(unionVar.memory)
}
func (unionVar *Cdr_u) fn() *compiled_fn {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*compiled_fn)(unionVar.memory)
}

type Object_s struct {
	car *Object_s
	cdr Cdr_u
}
type Object Object_s

var atomic Object
var p_atomic *Object_s = (*Object_s)(unsafe.Pointer(&atomic))
var compiled Object
var p_compiled *Object_s = (*Object_s)(unsafe.Pointer(&compiled))
var p_environment *Object_s = nil
var p_nil *Object_s = nil
var p_quote *Object_s = nil
var p_atom *Object_s = nil
var p_eq *Object_s = nil
var p_cons *Object_s = nil
var p_car *Object_s = nil
var p_cdr *Object_s = nil
var p_cond *Object_s = nil
var p_eval *Object_s = nil
var p_apply *Object_s = nil
var p_defglobal *Object_s = nil
var p_dot_symbol_dump *Object_s = nil

// nilp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:400
/* Forward references. */ //
/* Objects (lists, atoms, etc.). */ //
/* Head item in this list, unless a BUILTIN_CAR node. */ //
/* Tail list, unless car is a BUILTIN_CAR node. */ //
/* If CAR of an Object is not one of these, it is a List. */ //
/* Object.cdr is a struct Symbol_s *. */ //
/* Object.cdr is a compiled_fn. */ //
/* Current evaluation environment (list of bindings). */ //
/* Input/output as (). */ //
/* Also input as '. */ //
//
func nilp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int64(uintptr(unsafe.Pointer(list))) == int64(uintptr(unsafe.Pointer(p_nil)))])))
}

// atomp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:406
/* Whether object is an atom in the traditional Lisp sense */ //
//
func atomp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int32(int8((nilp(list)))) != 0 || int64(uintptr(unsafe.Pointer((*list).car))) == int64(uintptr(unsafe.Pointer(p_atomic)))])))
}

// atomicp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:412
/* Whether object is an atom in this implementation */ //
//
func atomicp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int8((noarch.NotInt8(nilp(list)))) != 0 && int64(uintptr(unsafe.Pointer((*list).car))) == int64(uintptr(unsafe.Pointer(p_atomic)))])))
}

// compiledp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:417
func compiledp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int8((noarch.NotInt8(atomp(list)))) != 0 && int64(uintptr(unsafe.Pointer((*list).car))) == int64(uintptr(unsafe.Pointer(p_compiled)))])))
}

// listp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:422
func listp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int8((noarch.NotInt8(atomp(list)))) != 0 && int8((noarch.NotInt8(compiledp(list)))) != 0])))
}

// finalp - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:427
func finalp(list *Object_s) int8 {
	return int8((int8(map[bool]int32{false: 0, true: 1}[int32(int8((listp(list)))) != 0 && int32(int8((nilp((*(*list).cdr.obj()))))) != 0])))
}

// list_car - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:432
func list_car(list *Object_s) *Object_s {
	func() {
		if int32(int8((listp(list)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(list)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(434)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return (*list).car
}

// list_cdr - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:438
func list_cdr(list *Object_s) *Object_s {
	func() {
		if int32(int8((listp(list)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(list)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(440)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return (*(*list).cdr.obj())
}

// object_symbol - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:444
func object_symbol(atom *Object_s) *Symbol_s {
	func() {
		if int32(int8((atomicp(atom)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("atomicp(atom)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(446)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return (*(*atom).cdr.sym())
}

// object_compiled - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:450
func object_compiled(compiled *Object_s) compiled_fn {
	func() {
		if int32(int8((compiledp(compiled)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("compiledp(compiled)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(452)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return compiled_fn((*(*compiled).cdr.fn()))
}

// object_new - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:456
func object_new(car *Object_s, cdr *Object_s) *Object_s {
	var obj *Object_s
	obj = (*Object_s)(noarch.Malloc(int32(16)))
	if obj == nil {
		for {
			var m *byte = (&[]byte("no more memory\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(998)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	allocations += 1
	allocations_total += int64(uint64(16))
	(*obj).car = car
	(*(*obj).cdr.obj()) = cdr
	return obj
}

// object_new_compiled - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:473
func object_new_compiled(fn compiled_fn) *Object_s {
	var obj *Object_s
	obj = (*Object_s)(noarch.Malloc(int32(16)))
	if obj == nil {
		for {
			var m *byte = (&[]byte("no more memory\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(998)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	allocations += 1
	allocations_total += int64(uint64(16))
	(*obj).car = p_compiled
	(*(*obj).cdr.fn()) = fn
	return obj
}

// symbol_new - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:498
/* Symbols. */ //
//
func symbol_new(name *byte) *Symbol_s {
	var sym *Symbol_s
	sym = (*Symbol_s)(noarch.Malloc(int32(8)))
	if sym == nil {
		for {
			var m *byte = (&[]byte("no more memory\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(998)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	allocations += 1
	allocations_total += int64(uint64(8))
	(*sym).name = name
	return sym
}

type map_sym_t struct {
	base map_base_t
	ref  *Symbol_s
	tmp  *Symbol_s
}

var map_sym map_sym_t

// symbol_lookup - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:519
/* Map of symbols (keys) to values. */ //
//
func symbol_lookup(name *byte) *Symbol_s {
	var p_sym *unsafe.Pointer = (*unsafe.Pointer)(unsafe.Pointer((func() *Symbol_s {
		(*(&map_sym)).ref = (*Symbol_s)(map_get_(&(*(&map_sym)).base, name))
		return (*(&map_sym)).ref
	}())))
	return (*Symbol_s)(func() unsafe.Pointer {
		if p_sym != nil {
			return *p_sym
		} else {
			return (nil)
		}
	}())
}

var symbol_strdup int8 = int8((int8(int32(1))))

// symbol_sym - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:528
func symbol_sym(name *byte) *Symbol_s {
	var sym *Symbol_s = symbol_lookup(name)
	if sym != nil {
		return sym
	}
	sym = symbol_new(func() *byte {
		if int32(int8((symbol_strdup))) != 0 {
			return string_duplicate(name)
		} else {
			return name
		}
	}())
	func() *Symbol_s {
		(*(&map_sym)).tmp = sym
		return (*(&map_sym)).tmp
	}()
	map_set_(&(*(&map_sym)).base, name, unsafe.Pointer(&(*(&map_sym)).tmp), int32(8))
	return sym
}

// symbol_dump - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:542
func symbol_dump() {
	var iter map_iter_t = map_iter_()
	var key *byte
	for (func() *byte {
		tempVar := map_next_(&(*(&map_sym)).base, &iter)
		key = tempVar
		return tempVar
	}()) != nil {
		noarch.Printf((&[]byte("%s -> %p\n\x00")[0]), key, symbol_lookup(key))
	}
}

var p_sym_t *Symbol_s = nil
var p_sym_quote *Symbol_s = nil

// binding_new - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:558
/* Environment (bindings). */ //
//
func binding_new(sym *Object_s, val *Object_s) *Object_s {
	func() {
		if int32(int8((atomicp(sym)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("atomicp(sym)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(560)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return object_new(sym, val)
}

// binding_lookup - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:571
/* Bindings; each binding is either an atom (meaning its symbol is
   explicitly unbound) or a key/value cons (the symbol is in the car,
   its binding is in the cdr). */ //
/* Originally this used a recursive algorithm, but tail-recursion
   optimization wasn't being done by gcc -g -O0, and it was annoying
   to find oneself inside such deep stack traces during debugging.
*/ //
//
func binding_lookup(what *byte, key *Symbol_s, bindings *Object_s) *Object_s {
	if int8((nilp(bindings))) != 0 {
		return p_nil
	}
	func() {
		if int32(int8((listp(bindings)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(bindings)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(576)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	for ; int8((noarch.NotInt8(nilp(bindings)))) != 0; bindings = list_cdr(bindings) {
		var binding *Object_s = list_car(bindings)
		if int32(int8((atomicp(binding)))) != 0 && int64(uintptr(unsafe.Pointer(object_symbol(binding)))) == int64(uintptr(unsafe.Pointer(key))) {
			return p_nil
		}
		{
			var symbol *Object_s = list_car(binding)
			if int32(int8((atomicp(symbol)))) != 0 && int64(uintptr(unsafe.Pointer(object_symbol(symbol)))) == int64(uintptr(unsafe.Pointer(key))) {
				return binding
			}
		}
	}
	return p_nil
}

var token_lookahead *byte
var lookahead_valid int32 = int32(0)

type buffer_s struct {
	size     size_t
	used     size_t
	contents *byte
}

// buffer_new - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:611
/* Input */ //
//
func buffer_new(initial_size size_t) *buffer_s {
	var buf *buffer_s = (*buffer_s)(noarch.Malloc(int32(24)))
	var contents *byte = (*byte)(noarch.Malloc(int32(uint32((size_t(initial_size))))))
	if buf == nil || contents == nil {
		for {
			var m *byte = (&[]byte("cannot allocate buffer for lexemes\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(998)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	(*buf).size = initial_size
	(*buf).used = size_t(int32(0))
	(*buf).contents = contents
	return buf
}

// buffer_append - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:626
/* TODO: realloc() */ //
//
func buffer_append(buf *buffer_s, ch byte) {
	if size_t((*buf).used) >= size_t((*buf).size) {
		for {
			var m *byte = (&[]byte("lexeme too long for this build\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(997)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	*((*byte)(func() unsafe.Pointer {
		tempVar := (*buf).contents
		return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((func() size_t {
			tempVar := &(*buf).used
			defer func() {
				*tempVar += 1
			}()
			return *tempVar
		}()))))*unsafe.Sizeof(*tempVar))
	}())) = ch
}

// buffer_tostring - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:633
func buffer_tostring(buf *buffer_s) *byte {
	buffer_append(buf, '\x00')
	return (*buf).contents
}

// token_putback - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:641
func token_putback(token *byte) {
	func() {
		if (map[bool]int32{false: 0, true: 1}[lookahead_valid == int32(0)]) != 0 {
		} else {
			linux.AssertFail((&[]byte("lookahead_valid == 0\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(643)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	token_lookahead = token
	lookahead_valid = int32(1)
}

var my_getc_next int32
var my_getc_file *noarch.File

// my_getc - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:651
func my_getc(input *noarch.File) int32 {
	if int64(uintptr(unsafe.Pointer(my_getc_file))) == int64(uintptr(unsafe.Pointer(input))) {
		my_getc_file = nil
		return my_getc_next
	}
	func() {
		if (map[bool]int32{false: 0, true: 1}[my_getc_file == nil || (&[]byte("No support for un-getting on two streams at the same time\x00")[0]) == nil]) != 0 {
		} else {
			linux.AssertFail((&[]byte("!my_getc_file || (\"No support for un-getting on two streams at the same time\" == NULL)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(659)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return noarch.Fgetc(input)
}

// my_ungetc - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:664
func my_ungetc(ch int32, input *noarch.File) {
	func() {
		if (map[bool]int32{false: 0, true: 1}[my_getc_file == nil || (&[]byte("No support for un-getting on two streams at the same time\x00")[0]) == nil]) != 0 {
		} else {
			linux.AssertFail((&[]byte("!my_getc_file || (\"No support for un-getting on two streams at the same time\" == NULL)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(667)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	my_getc_file = input
	my_getc_next = ch
}

// token_get - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:673
func token_get(input *noarch.File, buf *buffer_s) (c2goDefaultReturn *byte) {
	var ch int32
	func() size_t {
		(*(buf)).used = size_t(int32(0))
		return (*(buf)).used
	}()
	if lookahead_valid != 0 {
		lookahead_valid = int32(0)
		return token_lookahead
	}
	for {
		if (func() int32 {
			tempVar := my_getc(input)
			ch = tempVar
			return tempVar
		}()) == -int32(1) {
			for {
				var m *byte = nil
				if m != nil {
					noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
				}
				if int8((noarch.NotInt8(quiet))) != 0 {
					noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
				}
				noarch.Exit((int32(0)))
				if noarch.NotInt32((int32(0))) != 0 {
					break
				}
			}
		}
		if ch == int32(';') {
			for (func() int32 {
				tempVar := my_getc(input)
				ch = tempVar
				return tempVar
			}()) != -int32(1) && ch != int32('\n') {
			}
		}
		if noarch.NotInt32((int32(*((*uint16)(func() unsafe.Pointer {
			tempVar := (*linux.CtypeLoc())
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((ch))*unsafe.Sizeof(*tempVar))
		}()))) & int32(uint16(_ISspace)))) != 0 {
			break
		}
	}
	buffer_append(buf, byte(ch))
	if noarch.Strchr((&[]byte("()'\x00")[0]), ch) != nil {
		return buffer_tostring(buf)
	}
	for {
		if (func() int32 {
			tempVar := my_getc(input)
			ch = tempVar
			return tempVar
		}()) == -int32(1) {
			for {
				var m *byte = nil
				if m != nil {
					noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
				}
				if int8((noarch.NotInt8(quiet))) != 0 {
					noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
				}
				noarch.Exit((int32(0)))
				if noarch.NotInt32((int32(0))) != 0 {
					break
				}
			}
		}
		if noarch.Strchr((&[]byte("()'\x00")[0]), ch) != nil || int32(*((*uint16)(func() unsafe.Pointer {
			tempVar := (*linux.CtypeLoc())
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)((ch))*unsafe.Sizeof(*tempVar))
		}())))&int32(uint16(_ISspace)) != 0 {
			my_ungetc(ch, input)
			return buffer_tostring(buf)
		}
		buffer_append(buf, byte(ch))
	}
	return
}

// object_read - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:713
func object_read(input *noarch.File, buf *buffer_s) *Object_s {
	var token *byte
	token = token_get(input, buf)
	if noarch.NotInt32(noarch.Strcmp(token, (&[]byte("(\x00")[0]))) != 0 {
		return list_read(input, buf)
	}
	if noarch.NotInt32(noarch.Strcmp(token, (&[]byte("'\x00")[0]))) != 0 {
		var tmp *Object_s = object_read(input, buf)
		return object_new(object_new(p_atomic, (*Object_s)(unsafe.Pointer((p_sym_quote)))), object_new(tmp, p_nil))
	}
	if noarch.NotInt32(noarch.Strcmp(token, (&[]byte(")\x00")[0]))) != 0 {
		for {
			var m *byte = (&[]byte("unbalanced close paren\x00")[0])
			if m != nil {
				noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
			}
			if int8((noarch.NotInt8(quiet))) != 0 {
				noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
			}
			noarch.Exit((int32(1)))
			if noarch.NotInt32((int32(0))) != 0 {
				break
			}
		}
	}
	return object_new(p_atomic, (*Object_s)(unsafe.Pointer((symbol_sym(token)))))
}

// list_read - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:736
/* Make sure we first read the object before going on to read the rest of the list. */ //
//
func list_read(input *noarch.File, buf *buffer_s) *Object_s {
	var token *byte = token_get(input, buf)
	var tmp *Object_s
	if noarch.NotInt32(noarch.Strcmp(token, (&[]byte(")\x00")[0]))) != 0 {
		return p_nil
	}
	if noarch.NotInt32(noarch.Strcmp(token, (&[]byte(".\x00")[0]))) != 0 {
		tmp = object_read(input, buf)
		if noarch.Strcmp(token_get(input, buf), (&[]byte(")\x00")[0])) != 0 {
			for {
				var m *byte = (&[]byte("missing close parenthese for simple list\x00")[0])
				if m != nil {
					noarch.Fprintf(stderr, (&[]byte("%s\n\x00")[0]), m)
				}
				if int8((noarch.NotInt8(quiet))) != 0 {
					noarch.Fprintf(stderr, (&[]byte("allocations: %lld; total: %lld\n\x00")[0]), allocations, allocations_total)
				}
				noarch.Exit((int32(3)))
				if noarch.NotInt32((int32(0))) != 0 {
					break
				}
			}
		}
		return tmp
	}
	token_putback(token)
	tmp = object_read(input, buf)
	return object_new(tmp, list_read(input, buf))
}

// quotep - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:761
/* Output. */ //
/* true if object is (quote arg) */ //
/* TODO: Decide whether this look up quote in the current env to do the check */ //
//
func quotep(obj *Object_s) (c2goDefaultReturn int8) {
	if int8((noarch.NotInt8(listp(obj)))) != 0 || int8((noarch.NotInt8(finalp(list_cdr(obj))))) != 0 {
		return int8((int8(int32(0))))
	}
	{
		var car *Object_s = list_car(obj)
		return int8((int8(map[bool]int32{false: 0, true: 1}[int32(int8((compiledp(car)))) != 0 && int64(uintptr(noarch.CastInterfaceToPointer((func(*byte, *Object_s, *Object_s) *Object_s)((object_compiled(car)))))) == int64(uintptr(noarch.CastInterfaceToPointer(f_quote)))])))
	}
	return
}

// object_write - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:773
/* TODO: Print name of function. */ //
//
func object_write(output *noarch.File, obj *Object_s) {
	if int8((nilp(obj))) != 0 {
		noarch.Fprintf(output, (&[]byte("()\x00")[0]))
		return
	}
	if int8((atomicp(obj))) != 0 {
		noarch.Fprintf(output, (&[]byte("%s\x00")[0]), ((*(object_symbol(obj))).name))
		return
	}
	if int8((compiledp(obj))) != 0 {
		noarch.Fprintf(output, (&[]byte("*COMPILED*\x00")[0]))
		return
	}
	if int8((quotep(obj))) != 0 {
		noarch.Fprintf(output, (&[]byte("'\x00")[0]))
		object_write(output, list_car(list_cdr(obj)))
		return
	}
	noarch.Fprintf(output, (&[]byte("(\x00")[0]))
	for {
		object_write(output, list_car(obj))
		if int8((finalp(obj))) != 0 {
			noarch.Fprintf(output, (&[]byte(")\x00")[0]))
			return
		}
		obj = list_cdr(obj)
		if int8((noarch.NotInt8(listp(obj)))) != 0 {
			noarch.Fprintf(output, (&[]byte(" . \x00")[0]))
			object_write(output, obj)
			noarch.Fprintf(output, (&[]byte(")\x00")[0]))
			return
		}
		noarch.Fprintf(output, (&[]byte(" \x00")[0]))
	}
}

// binding_for - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:827
/* Evaluation */ //
/* TODO: Throw an exception etc. */ //
//
func binding_for(what *byte, sym *Symbol_s, env *Object_s) *Object_s {
	var tmp *Object_s
	tmp = binding_lookup(what, sym, env)
	if int8((nilp(tmp))) != 0 {
		noarch.Fprintf(stderr, (&[]byte("Unbound symbol \"%s\"\n\x00")[0]), ((*(sym)).name))
		func() {
			if (map[bool]int32{false: 0, true: 1}[(&[]byte("unbound symbol\x00")[0]) == nil]) != 0 {
			} else {
				linux.AssertFail((&[]byte("\"unbound symbol\" == NULL\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(836)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
	}
	return list_cdr(tmp)
}

// eval - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:846
/* Does not support traditional lambdas or labels; just the built-ins and
   our "unique" apply.  */ //
//
func eval(what *byte, exp *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	if int32(int8((nilp(exp)))) != 0 || int32(int8((compiledp(exp)))) != 0 {
		return exp
	}
	if int8((atomicp(exp))) != 0 {
		return binding_for(what, object_symbol(exp), env)
	}
	func() {
		if int32(int8((listp(exp)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(exp)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(854)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var func_ *Object_s = eval(what, list_car(exp), env)
		var forms *Object_s = list_cdr(exp)
		if int8((atomp(list_car(exp)))) != 0 {
			what = (*(object_symbol(list_car(exp)))).name
		}
		if int8((compiledp(func_))) != 0 {
			var fn compiled_fn
			fn = object_compiled(func_)
			return fn(what, forms, env)
		}
		return apply(what, func_, func_, forms, env)
	}
	return
}

// assert_zedbap - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:908
/* Need to solve these problems:

   (defun x (a) (progn (setq a 6) (a)))
   (setq a 5)
   (x a)

   It should return 5, but if defun is implemented naively, it
   might return 6.

   (defun y (a b c) (cons (a b c)))
   (setq a 5)
   (setq b 6)
   (setq c 7)
   (y c a b)

   That should return (7 5 6).

   Consider: Make the built-in form work as much like a compiled form
   as possible.  The compiled form is passed in an arglist and the
   current environment, allowing it to manipulate both to a high
   degree.  The compiler lexically binds those arguments to parameter
   names, so an instance of the compiled function doesn't put those
   parameter names into the environment.  So, either implement some
   form of lexical binding of names here, or consider a nameless
   solution such as having a particular form eval to the argument list
   and another to the environment.  This means explicitly supporting
   an eval() function that, unlike traditional Lisp, takes an env
   argument.
*/ //
/* mename */ //
/* formlistparamname */ //
/* envparamname */ //
/* envparamname */ //
/* body */ //
//
func assert_zedbap(zedba *Object_s) {
	func() {
		if int32(int8((listp(zedba)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(zedba)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(910)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((listp(list_car(zedba))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(list_car(zedba))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(912)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((atomicp(list_car(list_car(zedba)))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("atomicp(list_car(list_car(zedba)))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(913)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((atomicp(list_car(list_cdr(list_car(zedba))))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("atomicp(list_car(list_cdr(list_car(zedba))))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(914)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((atomicp(list_car(list_cdr(list_cdr(list_car(zedba)))))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("atomicp(list_car(list_cdr(list_cdr(list_car(zedba)))))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(915)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((finalp(list_cdr(list_cdr(list_car(zedba))))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(list_cdr(list_cdr(list_car(zedba))))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(916)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((finalp(list_cdr(zedba))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(list_cdr(zedba))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(918)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
}

// apply - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:934
/* Apply a zedba, which is an self/arglist/env macro version of the
   classic lambda.  The form of a zedba is

   ((mename formlistparamname envparamname) body)

   where body is to be evaluated with formlistparamname bound to the
   list of forms of the invocation, envparamname to the environment
   for the evaluation of the containing expression, and mename bound
   to the zedba itself (for easy recursive references).  Note that the
   caller might want to pass something else to be bound to mename, in
   case this proves useful (e.g. a limit on the # of recursive
   invocations could be implemented this way), so this is allowed.  */ //
//
func apply(what *byte, func_ *Object_s, me *Object_s, forms *Object_s, env *Object_s) *Object_s {
	var meparamname *Object_s
	var formlistparamname *Object_s
	var envparamname *Object_s
	assert_zedbap(func_)
	assert_zedbap(me)
	{
		var params *Object_s = list_car(func_)
		func() {
			if int32(int8((listp(params)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(params)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(946)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		meparamname = list_car(params)
		func() {
			if int32(int8((listp(list_cdr(params))))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(list_cdr(params))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(950)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		formlistparamname = list_car(list_cdr(params))
		func() {
			if int32(int8((finalp(list_cdr(list_cdr(params)))))) != 0 {
			} else {
				linux.AssertFail((&[]byte("finalp(list_cdr(list_cdr(params)))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(954)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		envparamname = list_car(list_cdr(list_cdr(params)))
	}
	return eval(what, list_car(list_cdr(func_)), (object_new(binding_new((meparamname), (func_)), (object_new(binding_new((formlistparamname), (forms)), (object_new(binding_new((envparamname), (env)), (env))))))))
}

// f_quote - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:967
/* (quote form) => form */ //
//
func f_quote(what *byte, args *Object_s, env *Object_s) *Object_s {
	func() {
		if int32(int8((finalp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(969)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return list_car(args)
}

// f_atom - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:975
/* (atom atom) => t if atom is an atom (including nil), nil otherwise */ //
//
func f_atom(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((finalp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(977)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var arg *Object_s = eval(what, list_car(args), env)
		return func() *Object_s {
			if int32(int8((atomp(arg)))) != 0 {
				return object_new(p_atomic, (*Object_s)(unsafe.Pointer((p_sym_t))))
			} else {
				return p_nil
			}
		}()
	}
	return
}

// f_eq - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:987
/* (eq left-atom right-atom) => t if args are equal, nil otherwise */ //
/* All nils are equal to each other in this implementation */ //
//
func f_eq(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(989)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((finalp(list_cdr(args))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(list_cdr(args))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(990)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var left *Object_s = eval(what, list_car(args), env)
		var right *Object_s = eval(what, list_car(list_cdr(args)), env)
		func() {
			if int32(int8((atomp(left)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("atomp(left)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(996)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		func() {
			if int32(int8((atomp(right)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("atomp(right)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(997)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		if int64(uintptr(unsafe.Pointer(left))) == int64(uintptr(unsafe.Pointer(right))) {
			return object_new(p_atomic, (*Object_s)(unsafe.Pointer((p_sym_t))))
		}
		if int32(int8((nilp(left)))) != 0 || int32(int8((nilp(right)))) != 0 {
			return p_nil
		}
		return func() *Object_s {
			if int64(uintptr(unsafe.Pointer(object_symbol(left)))) == int64(uintptr(unsafe.Pointer(object_symbol(right)))) {
				return object_new(p_atomic, (*Object_s)(unsafe.Pointer((p_sym_t))))
			} else {
				return p_nil
			}
		}()
	}
	return
}

// f_cons - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1010
/* (cons car-arg cdr-arg) => (car-arg cdr-arg) */ //
//
func f_cons(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1012)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if int32(int8((finalp(list_cdr(args))))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(list_cdr(args))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1013)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var car *Object_s = eval(what, list_car(args), env)
		var cdr *Object_s = eval(what, list_car(list_cdr(args)), env)
		return object_new(car, cdr)
	}
	return
}

// f_car - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1024
/* (car cons-arg) : cons-arg is a list => car of cons-arg */ //
//
func f_car(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((finalp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1026)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var arg *Object_s = eval(what, list_car(args), env)
		func() {
			if int32(int8((listp(arg)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(arg)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1031)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		return list_car(arg)
	}
	return
}

// f_cdr - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1038
/* (cdr cons-arg) : cons-arg is a list => cdr of cons-arg */ //
//
func f_cdr(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((finalp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("finalp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1040)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var arg *Object_s = eval(what, list_car(args), env)
		func() {
			if int32(int8((listp(arg)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(arg)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1045)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		return list_cdr(arg)
	}
	return
}

// f_cond - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1055
/* (cond ifthen-args ...) : each ifthen-args is an ifthen-pair; each
   ifthen-pair is a list of form (if-arg then-form) => eval(then-form)
   for the first if-arg in the list that is not nil (true), otherwise
   nil. */ //
//
func f_cond(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	if int8((nilp(args))) != 0 {
		return p_nil
	}
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1060)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var pair *Object_s = list_car(args)
		func() {
			if int32(int8((listp(pair)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(pair)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1065)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		func() {
			if int32(int8((finalp(list_cdr(pair))))) != 0 {
			} else {
				linux.AssertFail((&[]byte("finalp(list_cdr(pair))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1066)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		{
			var if_arg *Object_s = list_car(pair)
			var then_form *Object_s = list_car(list_cdr(pair))
			if int8((noarch.NotInt8(nilp(eval(what, if_arg, env))))) != 0 {
				return eval(what, then_form, env)
			}
			return f_cond(what, list_cdr(args), env)
		}
	}
	return
}

// f_defglobal - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1085
/* (defglobal) => global environment
   (defglobal newenv) => '() with newenv as the new global environment (SIDE EFFECT)
   (defglobal key value) => '() with new global environment prepended (via cons)
   			    with (key . value) (SIDE EFFECT)
*/ //
/* This form allows direct replacement of global environment.
   E.g. (defglobal (cdr (defglobal))) pops off the top binding. */ //
//
func f_defglobal(what *byte, args *Object_s, env *Object_s) *Object_s {
	if int8((nilp(args))) != 0 {
		return p_environment
	}
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1090)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	if int8((nilp(list_cdr(args)))) != 0 {
		p_environment = eval(what, list_car(args), env)
	} else {
		func() {
			if int32(int8((finalp(list_cdr(args))))) != 0 {
			} else {
				linux.AssertFail((&[]byte("finalp(list_cdr(args))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1098)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		{
			var sym *Object_s = eval(what, list_car(args), env)
			var form *Object_s = eval(what, list_car(list_cdr(args)), env)
			func() {
				if int32(int8((atomicp(sym)))) != 0 {
				} else {
					linux.AssertFail((&[]byte("atomicp(sym)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1104)), (&[]byte("void print_number(int *)\x00")[0]))
				}
			}()
			p_environment = object_new(binding_new((object_new(p_atomic, (*Object_s)(unsafe.Pointer((object_symbol(sym)))))), (form)), (p_environment))
		}
	}
	return p_nil
}

// f_eval - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1116
/* (eval arg [env]) => arg evaluated with respect to environment env (default is current env) */ //
//
func f_eval(what *byte, args *Object_s, env *Object_s) *Object_s {
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1118)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	func() {
		if (map[bool]int32{false: 0, true: 1}[int32(int8((nilp(list_cdr(args))))) != 0 || int32(int8((finalp(list_cdr(args))))) != 0]) != 0 {
		} else {
			linux.AssertFail((&[]byte("nilp(list_cdr(args)) || finalp(list_cdr(args))\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1119)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	return eval(what, eval(what, list_car(args), env), func() *Object_s {
		if int32(int8((nilp(list_cdr(args))))) != 0 {
			return env
		} else {
			return eval(what, list_car(list_cdr(args)), env)
		}
	}())
}

// f_apply - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1128
/* (apply zedba me forms [env]) => zedba invoked with reference to
   (presumably) itself, forms to be bound to zedba's arguments, and
   environment for such bindings (default is current env) */ //
//
func f_apply(what *byte, args *Object_s, env *Object_s) (c2goDefaultReturn *Object_s) {
	func() {
		if int32(int8((listp(args)))) != 0 {
		} else {
			linux.AssertFail((&[]byte("listp(args)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1130)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	{
		var func_ *Object_s = eval(what, list_car(args), env)
		var rest *Object_s = list_cdr(args)
		if int8((atomp(list_car(args)))) != 0 {
			what = (*(object_symbol(list_car(args)))).name
		}
		func() {
			if int32(int8((listp(rest)))) != 0 {
			} else {
				linux.AssertFail((&[]byte("listp(rest)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1141)), (&[]byte("void print_number(int *)\x00")[0]))
			}
		}()
		{
			var me *Object_s = eval(what, list_car(rest), env)
			var new_rest *Object_s = list_cdr(rest)
			rest = new_rest
			func() {
				if int32(int8((listp(rest)))) != 0 {
				} else {
					linux.AssertFail((&[]byte("listp(rest)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1148)), (&[]byte("void print_number(int *)\x00")[0]))
				}
			}()
			{
				var forms *Object_s = eval(what, list_car(rest), env)
				var new_rest *Object_s = list_cdr(rest)
				rest = new_rest
				func() {
					if (map[bool]int32{false: 0, true: 1}[int32(int8((nilp(rest)))) != 0 || int32(int8((finalp(rest)))) != 0]) != 0 {
					} else {
						linux.AssertFail((&[]byte("nilp(rest) || finalp(rest)\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1155)), (&[]byte("void print_number(int *)\x00")[0]))
					}
				}()
				return apply(what, func_, me, forms, func() *Object_s {
					if int32(int8((nilp(rest)))) != 0 {
						return p_nil
					} else {
						return eval(what, list_car(rest), env)
					}
				}())
			}
		}
	}
	return
}

// f_dot_symbol_dump - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1164
/* (.symbol_dump) : dump symbol names along with their struct Symbol_s * objects */ //
//
func f_dot_symbol_dump(what *byte, args *Object_s, env *Object_s) *Object_s {
	func() {
		if (map[bool]int32{false: 0, true: 1}[args == nil]) != 0 {
		} else {
			linux.AssertFail((&[]byte("!args\x00")[0]), (&[]byte("/home/craig/github/LispZero/lisp-zero-single.c\x00")[0]), uint32(int32(1166)), (&[]byte("void print_number(int *)\x00")[0]))
		}
	}()
	symbol_dump()
	return p_nil
}

// initialize_builtin - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1173
func initialize_builtin(sym *byte, fn compiled_fn) *Object_s {
	var tmp *Object_s
	p_environment = object_new(binding_new((object_new(p_atomic, (*Object_s)(unsafe.Pointer((symbol_sym(sym)))))), (func() *Object_s {
		tmp = object_new_compiled(compiled_fn(fn))
		return tmp
	}())), (p_environment))
	return tmp
}

// initialize - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1184
/* TODO: Decide on a better name. */ //
//
func initialize() {
	p_sym_t = symbol_sym((&[]byte("t\x00")[0]))
	p_sym_quote = symbol_sym((&[]byte("quote\x00")[0]))
	symbol_strdup = int8((int8(int32(0))))
	p_quote = initialize_builtin((&[]byte("quote\x00")[0]), f_quote)
	p_atom = initialize_builtin((&[]byte("atom\x00")[0]), f_atom)
	p_eq = initialize_builtin((&[]byte("eq\x00")[0]), f_eq)
	p_cons = initialize_builtin((&[]byte("cons\x00")[0]), f_cons)
	p_car = initialize_builtin((&[]byte("car\x00")[0]), f_car)
	p_cdr = initialize_builtin((&[]byte("cdr\x00")[0]), f_cdr)
	p_cond = initialize_builtin((&[]byte("cond\x00")[0]), f_cond)
	p_eval = initialize_builtin((&[]byte("eval\x00")[0]), f_eval)
	p_apply = initialize_builtin((&[]byte("apply\x00")[0]), f_apply)
	p_defglobal = initialize_builtin((&[]byte("defglobal\x00")[0]), f_defglobal)
	p_dot_symbol_dump = initialize_builtin((&[]byte(".symbol_dump\x00")[0]), f_dot_symbol_dump)
	symbol_strdup = int8((int8(int32(1))))
}

// main - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1212
func main() {
	argc := int32(len(os.Args))
	argv__multiarray := [][]byte{}
	argv__array := []*byte{}
	for _, argvSingle := range os.Args {
		argv__multiarray = append(argv__multiarray, append([]byte(argvSingle), 0))
	}
	for _, argvSingle := range argv__multiarray {
		argv__array = append(argv__array, &argvSingle[0])
	}
	argv := *(***byte)(unsafe.Pointer(&argv__array))
	if argc > int32(1) && noarch.NotInt32(noarch.Strcmp(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(int32(1))*unsafe.Sizeof(*argv)))), (&[]byte("-q\x00")[0]))) != 0 {
		quiet = int8((int8(int32(1))))
	}
	noarch.Memset(unsafe.Pointer(&map_sym), int32(0), int32(32))
	var buf *buffer_s
	initialize()
	buf = buffer_new(size_t(int32(1000)))
	for {
		var obj *Object_s = object_read(stdin, buf)
		if int8((noarch.NotInt8(quiet))) != 0 {
			object_write(stdout, eval((&[]byte("stdin\x00")[0]), obj, p_environment))
			noarch.Fprintf(stdout, (&[]byte("\n\x00")[0]))
			noarch.Fflush(stdout)
		}
	}
	return
}

// debug_output - transpiled function from  /home/craig/github/LispZero/lisp-zero-single.c:1235
func debug_output(obj *Object_s) {
	object_write(stdout, obj)
	noarch.Fprintf(stdout, (&[]byte("\n\x00")[0]))
	noarch.Fflush(stdout)
}
func init() {
	stdin = noarch.Stdin
	stdout = noarch.Stdout
	stderr = noarch.Stderr
}

type _Bool int8
