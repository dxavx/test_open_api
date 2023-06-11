package handlers

import (
	"fmt"
	"net/http"
	"os"
	"test-open-api/internal"
	"test-open-api/internal/api"
)

func (s *Server) PostLogo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var request api.CreateLogoRequest

	if err := s.readRequest(r, &request); err != nil {
		s.replayError(ctx, w, &internal.Error{
			Code:       internal.DecodingError,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	err := saveToFile(*request.Filename, request.Logo)
	fmt.Println(err)

	s.writeResponse(ctx, w, http.StatusOK, api.CreateLogoResponses{
		Url: request.Filename,
	})

}

//func saveBase64ToFile(filename string, data string) error {
//	dec, err := base64.StdEncoding.DecodeString(data)
//	if err != nil {
//
//		fmt.Println("decoder error")
//		return err
//	}
//
//	fmt.Println("DECODER", dec)
//	fmt.Println()
//
//	f, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//
//	defer f.Close()
//
//	err = ioutil.WriteFile(filename, dec, 0644)
//	return err
//}

func saveToFile(filename string, data *[]byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {

		}
	}(f)

	err = os.WriteFile(filename, *data, 0644)
	return err
}
