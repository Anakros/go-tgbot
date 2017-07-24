// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/sticker.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Sticker) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Sticker) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Emoji) != 0 {
		buf.WriteString(`"emoji":`)
		fflib.WriteJsonString(buf, string(mj.Emoji))
		buf.WriteByte(',')
	}
	if len(mj.FileID) != 0 {
		buf.WriteString(`"file_id":`)
		fflib.WriteJsonString(buf, string(mj.FileID))
		buf.WriteByte(',')
	}
	if mj.FileSize != 0 {
		buf.WriteString(`"file_size":`)
		fflib.FormatBits2(buf, uint64(mj.FileSize), 10, mj.FileSize < 0)
		buf.WriteByte(',')
	}
	if mj.Height != 0 {
		buf.WriteString(`"height":`)
		fflib.FormatBits2(buf, uint64(mj.Height), 10, mj.Height < 0)
		buf.WriteByte(',')
	}
	if mj.MaskPosition != nil {
		if true {
			buf.WriteString(`"mask_position":`)

			{

				err = mj.MaskPosition.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.SetName) != 0 {
		buf.WriteString(`"set_name":`)
		fflib.WriteJsonString(buf, string(mj.SetName))
		buf.WriteByte(',')
	}
	if mj.Thumb != nil {
		if true {
			buf.WriteString(`"thumb":`)

			{

				err = mj.Thumb.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.Width != 0 {
		buf.WriteString(`"width":`)
		fflib.FormatBits2(buf, uint64(mj.Width), 10, mj.Width < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Stickerbase = iota
	ffj_t_Stickerno_such_key

	ffj_t_Sticker_Emoji

	ffj_t_Sticker_FileID

	ffj_t_Sticker_FileSize

	ffj_t_Sticker_Height

	ffj_t_Sticker_MaskPosition

	ffj_t_Sticker_SetName

	ffj_t_Sticker_Thumb

	ffj_t_Sticker_Width
)

var ffj_key_Sticker_Emoji = []byte("emoji")

var ffj_key_Sticker_FileID = []byte("file_id")

var ffj_key_Sticker_FileSize = []byte("file_size")

var ffj_key_Sticker_Height = []byte("height")

var ffj_key_Sticker_MaskPosition = []byte("mask_position")

var ffj_key_Sticker_SetName = []byte("set_name")

var ffj_key_Sticker_Thumb = []byte("thumb")

var ffj_key_Sticker_Width = []byte("width")

func (uj *Sticker) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Sticker) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Stickerbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Stickerno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffj_key_Sticker_Emoji, kn) {
						currentKey = ffj_t_Sticker_Emoji
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffj_key_Sticker_FileID, kn) {
						currentKey = ffj_t_Sticker_FileID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Sticker_FileSize, kn) {
						currentKey = ffj_t_Sticker_FileSize
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Sticker_Height, kn) {
						currentKey = ffj_t_Sticker_Height
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Sticker_MaskPosition, kn) {
						currentKey = ffj_t_Sticker_MaskPosition
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Sticker_SetName, kn) {
						currentKey = ffj_t_Sticker_SetName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Sticker_Thumb, kn) {
						currentKey = ffj_t_Sticker_Thumb
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Sticker_Width, kn) {
						currentKey = ffj_t_Sticker_Width
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Sticker_Width, kn) {
					currentKey = ffj_t_Sticker_Width
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Sticker_Thumb, kn) {
					currentKey = ffj_t_Sticker_Thumb
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Sticker_SetName, kn) {
					currentKey = ffj_t_Sticker_SetName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Sticker_MaskPosition, kn) {
					currentKey = ffj_t_Sticker_MaskPosition
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Sticker_Height, kn) {
					currentKey = ffj_t_Sticker_Height
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Sticker_FileSize, kn) {
					currentKey = ffj_t_Sticker_FileSize
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_Sticker_FileID, kn) {
					currentKey = ffj_t_Sticker_FileID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Sticker_Emoji, kn) {
					currentKey = ffj_t_Sticker_Emoji
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Stickerno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Sticker_Emoji:
					goto handle_Emoji

				case ffj_t_Sticker_FileID:
					goto handle_FileID

				case ffj_t_Sticker_FileSize:
					goto handle_FileSize

				case ffj_t_Sticker_Height:
					goto handle_Height

				case ffj_t_Sticker_MaskPosition:
					goto handle_MaskPosition

				case ffj_t_Sticker_SetName:
					goto handle_SetName

				case ffj_t_Sticker_Thumb:
					goto handle_Thumb

				case ffj_t_Sticker_Width:
					goto handle_Width

				case ffj_t_Stickerno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Emoji:

	/* handler: uj.Emoji type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Emoji = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FileID:

	/* handler: uj.FileID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.FileID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FileSize:

	/* handler: uj.FileSize type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.FileSize = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Height:

	/* handler: uj.Height type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Height = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaskPosition:

	/* handler: uj.MaskPosition type=models.MaskPosition kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.MaskPosition = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.MaskPosition == nil {
			uj.MaskPosition = new(MaskPosition)
		}

		err = uj.MaskPosition.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SetName:

	/* handler: uj.SetName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.SetName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Thumb:

	/* handler: uj.Thumb type=models.PhotoSize kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Thumb = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Thumb == nil {
			uj.Thumb = new(PhotoSize)
		}

		err = uj.Thumb.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Width:

	/* handler: uj.Width type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Width = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}