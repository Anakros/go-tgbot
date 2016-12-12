// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/game.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Game) MarshalJSON() ([]byte, error) {
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
func (mj *Game) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if mj.Animation != nil {
		if true {
			buf.WriteString(`"animation":`)

			{

				err = mj.Animation.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.Description != nil {
		buf.WriteString(`"description":`)
		fflib.WriteJsonString(buf, string(*mj.Description))
	} else {
		buf.WriteString(`"description":null`)
	}
	buf.WriteString(`,"photo":`)
	if mj.Photo != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Photo {
			if i != 0 {
				buf.WriteString(`,`)
			}
			if v != nil {
				/* Struct fall back. type=models.PhotoSize kind=struct */
				err = buf.Encode(&v)
				if err != nil {
					return err
				}
			} else {
				buf.WriteString(`null`)
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if len(mj.Text) != 0 {
		buf.WriteString(`"text":`)
		fflib.WriteJsonString(buf, string(mj.Text))
		buf.WriteByte(',')
	}
	buf.WriteString(`"text_entities":`)
	if mj.TextEntities != nil {
		buf.WriteString(`[`)
		for i, v := range mj.TextEntities {
			if i != 0 {
				buf.WriteString(`,`)
			}
			if v != nil {
				/* Struct fall back. type=models.MessageEntity kind=struct */
				err = buf.Encode(&v)
				if err != nil {
					return err
				}
			} else {
				buf.WriteString(`null`)
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	if mj.Title != nil {
		buf.WriteString(`,"title":`)
		fflib.WriteJsonString(buf, string(*mj.Title))
	} else {
		buf.WriteString(`,"title":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Gamebase = iota
	ffj_t_Gameno_such_key

	ffj_t_Game_Animation

	ffj_t_Game_Description

	ffj_t_Game_Photo

	ffj_t_Game_Text

	ffj_t_Game_TextEntities

	ffj_t_Game_Title
)

var ffj_key_Game_Animation = []byte("animation")

var ffj_key_Game_Description = []byte("description")

var ffj_key_Game_Photo = []byte("photo")

var ffj_key_Game_Text = []byte("text")

var ffj_key_Game_TextEntities = []byte("text_entities")

var ffj_key_Game_Title = []byte("title")

func (uj *Game) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Game) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Gamebase
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
				currentKey = ffj_t_Gameno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Game_Animation, kn) {
						currentKey = ffj_t_Game_Animation
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Game_Description, kn) {
						currentKey = ffj_t_Game_Description
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Game_Photo, kn) {
						currentKey = ffj_t_Game_Photo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Game_Text, kn) {
						currentKey = ffj_t_Game_Text
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Game_TextEntities, kn) {
						currentKey = ffj_t_Game_TextEntities
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Game_Title, kn) {
						currentKey = ffj_t_Game_Title
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Game_Title, kn) {
					currentKey = ffj_t_Game_Title
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Game_TextEntities, kn) {
					currentKey = ffj_t_Game_TextEntities
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Game_Text, kn) {
					currentKey = ffj_t_Game_Text
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Game_Photo, kn) {
					currentKey = ffj_t_Game_Photo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Game_Description, kn) {
					currentKey = ffj_t_Game_Description
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Game_Animation, kn) {
					currentKey = ffj_t_Game_Animation
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Gameno_such_key
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

				case ffj_t_Game_Animation:
					goto handle_Animation

				case ffj_t_Game_Description:
					goto handle_Description

				case ffj_t_Game_Photo:
					goto handle_Photo

				case ffj_t_Game_Text:
					goto handle_Text

				case ffj_t_Game_TextEntities:
					goto handle_TextEntities

				case ffj_t_Game_Title:
					goto handle_Title

				case ffj_t_Gameno_such_key:
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

handle_Animation:

	/* handler: uj.Animation type=models.Animation kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Animation = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Animation == nil {
			uj.Animation = new(Animation)
		}

		err = uj.Animation.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Description:

	/* handler: uj.Description type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Description = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Description = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Photo:

	/* handler: uj.Photo type=[]*models.PhotoSize kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Photo = nil
		} else {

			uj.Photo = []*PhotoSize{}

			wantVal := true

			for {

				var tmp_uj__Photo *PhotoSize

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Photo type=*models.PhotoSize kind=ptr quoted=false*/

				{

					if tok == fflib.FFTok_null {
						tmp_uj__Photo = nil
					} else {
						if tmp_uj__Photo == nil {
							tmp_uj__Photo = new(PhotoSize)
						}

						/* handler: tmp_uj__Photo type=models.PhotoSize kind=struct quoted=false*/

						{
							/* Falling back. type=models.PhotoSize kind=struct */
							tbuf, err := fs.CaptureField(tok)
							if err != nil {
								return fs.WrapErr(err)
							}

							err = json.Unmarshal(tbuf, &tmp_uj__Photo)
							if err != nil {
								return fs.WrapErr(err)
							}
						}

					}
				}

				uj.Photo = append(uj.Photo, tmp_uj__Photo)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Text:

	/* handler: uj.Text type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Text = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TextEntities:

	/* handler: uj.TextEntities type=[]*models.MessageEntity kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.TextEntities = nil
		} else {

			uj.TextEntities = []*MessageEntity{}

			wantVal := true

			for {

				var tmp_uj__TextEntities *MessageEntity

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__TextEntities type=*models.MessageEntity kind=ptr quoted=false*/

				{

					if tok == fflib.FFTok_null {
						tmp_uj__TextEntities = nil
					} else {
						if tmp_uj__TextEntities == nil {
							tmp_uj__TextEntities = new(MessageEntity)
						}

						/* handler: tmp_uj__TextEntities type=models.MessageEntity kind=struct quoted=false*/

						{
							/* Falling back. type=models.MessageEntity kind=struct */
							tbuf, err := fs.CaptureField(tok)
							if err != nil {
								return fs.WrapErr(err)
							}

							err = json.Unmarshal(tbuf, &tmp_uj__TextEntities)
							if err != nil {
								return fs.WrapErr(err)
							}
						}

					}
				}

				uj.TextEntities = append(uj.TextEntities, tmp_uj__TextEntities)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Title:

	/* handler: uj.Title type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Title = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Title = &tval

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