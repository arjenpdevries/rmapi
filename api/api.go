package api

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/juruen/rmapi/api/sync10"
	"github.com/juruen/rmapi/filetree"
	"github.com/juruen/rmapi/model"
	"github.com/juruen/rmapi/transport"
)

type ApiCtx interface {
	Filetree() *filetree.FileTreeCtx
	FetchDocument(docId, dstPath string) error
	CreateDir(parentId, name string) (model.Document, error)
	UploadDocument(parentId string, sourceDocPath string) (*model.Document, error)
	MoveEntry(src, dstDir *model.Node, name string) (*model.Node, error)
	DeleteEntry(node *model.Node) error
	Nuke() error
}
type UserToken struct {
	Scopes string
	*jwt.StandardClaims
}

// CreateApiCtx initializes an instance of ApiCtx
func CreateApiCtx(http *transport.HttpClientCtx) (ApiCtx, error) {
	userToken := http.Tokens.UserToken
	claims := UserToken{}
	jwt.ParseWithClaims(userToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	fmt.Println(claims.Scopes)
	fld := strings.Fields(claims.Scopes)
	isSync15 := false
	for _, f := range fld {
		switch f {
		case "sync:fox":
			fallthrough
		case "sync:tortois":
			fallthrough
		case "sync:hare":
			fmt.Println("New sync 1.5 not supported yet")
			isSync15 = true
			break
		case "sync:default":
			break
		}
	}
	if isSync15 {
		return sync10.CreateCtx(http)
	} else {
		return sync10.CreateCtx(http)
	}

}
