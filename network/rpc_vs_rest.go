package network

// restful api is known for being stateless and when you re using sessions its not restfull
// rpc has internal intelligence while rest needs external intelligence 
// because in rpc you re calling a well-known method but in rest there is no such a thing
// rpc is in transport layer while rest is in application layer
// being http doesnt clarify anything because you can use both of them through http
// http is a rest implemention but not every restful arch is working through http
// rpc hasnt any standards and not implemented for server/client model
// because there is no control while in rest there are headers and ...