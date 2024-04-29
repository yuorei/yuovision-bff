package lib

import "io"

func ReadSeekerToBytes(reader io.ReadSeeker) ([]byte, error) {
	// io.ReadSeekerのサイズを取得
	size, err := reader.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	// io.ReadSeekerの読み込み位置を先頭に戻す
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// io.ReadSeekerから[]byteに読み込み
	buffer := make([]byte, size)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
