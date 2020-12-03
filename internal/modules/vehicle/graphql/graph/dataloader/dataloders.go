/*
@Time : 2020/11/6 9:49
@Author : lai
@Description :
@File : dataloders
*/

//go:generate go run github.com/vektah/dataloaden UserLoader int *lai.com/gqlgen_study/demo02/graph/model.User
package dataloader

//type ctxKeyType struct{ name string }
//
//var ctxKey = ctxKeyType{"userCtx"}

//const loadersKey = "dataloaders"
//
//type Loaders struct {
//	UserById *UserLoader
//}
//
//func Middleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
//			UserById: &UserLoader{
//				maxBatch: 100,
//				wait:     1 * time.Millisecond,
//				fetch: func(ids []int) ([]*model.User, []error) {
//					placeholders := make([]string, len(ids))
//					args := make([]interface{}, len(ids))
//					for i := 0; i < len(ids); i++ {
//						placeholders[i] = "?"
//						args[i] = i
//					}
//					userById := map[int]*model.User{}
//					user := model.User{
//						ID:   1,
//						Name: "lai",
//					}
//
//					userById[user.ID] = &user
//
//					users := make([]*model.User, len(ids))
//					for i, id := range ids {
//						users[i] = userById[id]
//						i++
//					}
//
//					return users, nil
//				},
//			},
//		})
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}
//
//func For(ctx context.Context) *Loaders {
//	return ctx.Value(loadersKey).(*Loaders)
//}
