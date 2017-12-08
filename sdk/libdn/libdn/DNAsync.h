#pragma once

template <class T>
class DNAsync
{
	//private:
	//DNAsync() {}
public:
	// waits (synchronously, blocking) for the operation to complete
	virtual T* Wait() = 0;

	// waits, with a timeout (in msec)
	virtual T* Wait(unsigned int timeout) = 0;

	// returns whether the operation has completed
	virtual bool HasCompleted() = 0;

	// returns the operation result
	virtual T* GetResult() = 0;

	// sets a callback function to be called when the task completes
	virtual void SetCallback(void(__cdecl* callback)(DNAsync<T>*), void* userData) = 0;

	// sets a timeout callback function
	virtual void SetTimeoutCallback(void(__cdecl* callback)(DNAsync<T>*), unsigned int timeout) = 0;

	// gets user data from callback
	virtual void* GetUserData() = 0;

	// frees the result and the Async object itself
	virtual void Free() = 0;
};