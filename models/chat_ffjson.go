// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/chat.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Chat) MarshalJSON() ([]byte, error) {
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
func (mj *Chat) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.AllMembersAreAdministrators != false {
		if mj.AllMembersAreAdministrators {
			buf.WriteString(`"all_members_are_administrators":true`)
		} else {
			buf.WriteString(`"all_members_are_administrators":false`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Description) != 0 {
		buf.WriteString(`"description":`)
		fflib.WriteJsonString(buf, string(mj.Description))
		buf.WriteByte(',')
	}
	if len(mj.FirstName) != 0 {
		buf.WriteString(`"first_name":`)
		fflib.WriteJsonString(buf, string(mj.FirstName))
		buf.WriteByte(',')
	}
	buf.WriteString(`"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, mj.ID < 0)
	buf.WriteByte(',')
	if len(mj.InviteLink) != 0 {
		buf.WriteString(`"invite_link":`)
		fflib.WriteJsonString(buf, string(mj.InviteLink))
		buf.WriteByte(',')
	}
	if len(mj.LastName) != 0 {
		buf.WriteString(`"last_name":`)
		fflib.WriteJsonString(buf, string(mj.LastName))
		buf.WriteByte(',')
	}
	if mj.Photo != nil {
		if true {
			/* Struct fall back. type=models.ChatPhoto kind=struct */
			buf.WriteString(`"photo":`)
			err = buf.Encode(mj.Photo)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Title) != 0 {
		buf.WriteString(`"title":`)
		fflib.WriteJsonString(buf, string(mj.Title))
		buf.WriteByte(',')
	}
	if mj.Type != nil {
		buf.WriteString(`"type":`)
		fflib.WriteJsonString(buf, string(*mj.Type))
	} else {
		buf.WriteString(`"type":null`)
	}
	buf.WriteByte(',')
	if len(mj.Username) != 0 {
		buf.WriteString(`"username":`)
		fflib.WriteJsonString(buf, string(mj.Username))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Chatbase = iota
	ffj_t_Chatno_such_key

	ffj_t_Chat_AllMembersAreAdministrators

	ffj_t_Chat_Description

	ffj_t_Chat_FirstName

	ffj_t_Chat_ID

	ffj_t_Chat_InviteLink

	ffj_t_Chat_LastName

	ffj_t_Chat_Photo

	ffj_t_Chat_Title

	ffj_t_Chat_Type

	ffj_t_Chat_Username
)

var ffj_key_Chat_AllMembersAreAdministrators = []byte("all_members_are_administrators")

var ffj_key_Chat_Description = []byte("description")

var ffj_key_Chat_FirstName = []byte("first_name")

var ffj_key_Chat_ID = []byte("id")

var ffj_key_Chat_InviteLink = []byte("invite_link")

var ffj_key_Chat_LastName = []byte("last_name")

var ffj_key_Chat_Photo = []byte("photo")

var ffj_key_Chat_Title = []byte("title")

var ffj_key_Chat_Type = []byte("type")

var ffj_key_Chat_Username = []byte("username")

func (uj *Chat) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Chat) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Chatbase
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
				currentKey = ffj_t_Chatno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Chat_AllMembersAreAdministrators, kn) {
						currentKey = ffj_t_Chat_AllMembersAreAdministrators
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Chat_Description, kn) {
						currentKey = ffj_t_Chat_Description
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffj_key_Chat_FirstName, kn) {
						currentKey = ffj_t_Chat_FirstName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Chat_ID, kn) {
						currentKey = ffj_t_Chat_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Chat_InviteLink, kn) {
						currentKey = ffj_t_Chat_InviteLink
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_Chat_LastName, kn) {
						currentKey = ffj_t_Chat_LastName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Chat_Photo, kn) {
						currentKey = ffj_t_Chat_Photo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Chat_Title, kn) {
						currentKey = ffj_t_Chat_Title
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Chat_Type, kn) {
						currentKey = ffj_t_Chat_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_Chat_Username, kn) {
						currentKey = ffj_t_Chat_Username
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Chat_Username, kn) {
					currentKey = ffj_t_Chat_Username
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Chat_Type, kn) {
					currentKey = ffj_t_Chat_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Chat_Title, kn) {
					currentKey = ffj_t_Chat_Title
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Chat_Photo, kn) {
					currentKey = ffj_t_Chat_Photo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Chat_LastName, kn) {
					currentKey = ffj_t_Chat_LastName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Chat_InviteLink, kn) {
					currentKey = ffj_t_Chat_InviteLink
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Chat_ID, kn) {
					currentKey = ffj_t_Chat_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Chat_FirstName, kn) {
					currentKey = ffj_t_Chat_FirstName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Chat_Description, kn) {
					currentKey = ffj_t_Chat_Description
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Chat_AllMembersAreAdministrators, kn) {
					currentKey = ffj_t_Chat_AllMembersAreAdministrators
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Chatno_such_key
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

				case ffj_t_Chat_AllMembersAreAdministrators:
					goto handle_AllMembersAreAdministrators

				case ffj_t_Chat_Description:
					goto handle_Description

				case ffj_t_Chat_FirstName:
					goto handle_FirstName

				case ffj_t_Chat_ID:
					goto handle_ID

				case ffj_t_Chat_InviteLink:
					goto handle_InviteLink

				case ffj_t_Chat_LastName:
					goto handle_LastName

				case ffj_t_Chat_Photo:
					goto handle_Photo

				case ffj_t_Chat_Title:
					goto handle_Title

				case ffj_t_Chat_Type:
					goto handle_Type

				case ffj_t_Chat_Username:
					goto handle_Username

				case ffj_t_Chatno_such_key:
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

handle_AllMembersAreAdministrators:

	/* handler: uj.AllMembersAreAdministrators type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.AllMembersAreAdministrators = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.AllMembersAreAdministrators = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
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

		} else {

			outBuf := fs.Output.Bytes()

			uj.Description = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FirstName:

	/* handler: uj.FirstName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.FirstName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: uj.ID type=int64 kind=int64 quoted=false*/

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

			uj.ID = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InviteLink:

	/* handler: uj.InviteLink type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.InviteLink = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastName:

	/* handler: uj.LastName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.LastName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Photo:

	/* handler: uj.Photo type=models.ChatPhoto kind=struct quoted=false*/

	{
		/* Falling back. type=models.ChatPhoto kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Photo)
		if err != nil {
			return fs.WrapErr(err)
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

		} else {

			outBuf := fs.Output.Bytes()

			uj.Title = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Type = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Type = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Username:

	/* handler: uj.Username type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Username = string(string(outBuf))

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