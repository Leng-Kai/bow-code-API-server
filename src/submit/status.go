package submit

var Status2Flag = [32]uint{
	0x8000,		//	 0	Undefined:					Should not be found in response
	0x8000,		//	 1	In Queue:					Should not be found in response
	0x8000,		//	 2	Processing:					Should not be found in response
	0x0000,		//	 3	Accepted
	0x0010,		//	 4	Wrong Answer
	0x0008,		//	 5 	Time Limit Exceeded
	0x2000,		//	 6 	Compilation Error
	0x1100,		//	 7	Runtime Error (SIGSEGV)
	0x1200,		//	 8 	Runtime Error (SIGXFSZ)
	0x1300,		//	 9	Runtime Error (SIGFPE)
	0x1400,		//	10	Runtime Error (SIGABRT)
	0x1500,		//	11	Runtime Error (NZEC)
	0x1600,		//	12	Runtime Error (Other)
	0x8000,		//	13	Internal Error
	0x8000,		//	14 	Exec Format Error

	/* Not yet implemented */

	0x0004,		//		Memory Limit Exceed
}