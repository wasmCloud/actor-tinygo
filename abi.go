package actor

func HostCall(binding, namespace, operation string, payload []byte) ([]byte, error) {
	result := hostCall(
		stringToPointer(binding), uint32(len(binding)),
		stringToPointer(namespace), uint32(len(namespace)),
		stringToPointer(operation), uint32(len(operation)),
		bytesToPointer(payload), uint32(len(payload)),
	)
	if !result {
		errorLen := hostErrorLen()
		message := make([]byte, errorLen) // alloc
		hostError(bytesToPointer(message))

		return nil, &HostError{message: string(message)} // alloc
	}

	responseLen := hostResponseLen()
	response := make([]byte, responseLen) // alloc
	hostResponse(bytesToPointer(response))

	return response, nil
}

func (e *HostError) Error() string {
	return "Host error: " + e.message // alloc
}
