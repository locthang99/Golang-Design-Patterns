package proxy

type IShotHandle interface {
	HandleShotRequest(boss string)
}
