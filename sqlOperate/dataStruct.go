package sqlOperate

type User struct {
	Name     string
	Account  string
	Password string
	Sfp      string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetAccount() string {
	return u.Account
}

func (u *User) GetSfp() string {
	return u.Sfp
}

//错误消息处理
//==========================================================
type Wrongs struct {
	W string //错误信息
	S string //解决方法
	Wa string //返回的界面
}

func (w *Wrongs) Wrong() string{
	return w.W
}

func (w *Wrongs) Solustion() string{
	return w.S
}

func (w *Wrongs) Way() string{
	return w.Wa
}



//=======================================================================
type Notice struct {
	Title   string
	Content string
	Date    string
}

func (notice * Notice)GetTitle() string{
	return  notice.Title
}

func (notice *Notice)GetContent() string{
	return notice.Content
}

func (notice *Notice)GetDate() string{
	return notice.Date
}


//===============================================================================
type FileInformation struct{
	Title      string
	Date       string
	Path       string
	Tag        string
}

func (file *FileInformation)GetTag() string{
	return  file.Tag
}

func (file *FileInformation)GetDate() string{
	return file.Date
}

func (file *FileInformation)GetPath() string{
	return file.Path
}

func (file *FileInformation)GetTitle() string{
	return file.Title
}

//============================================================================
type feedbackItem struct {

	Time string

	Message string

	Account string
}

func (item *feedbackItem) GetTime() string{
	return item.Time
}

func (item *feedbackItem) GetMessage() string{
	return item.Message
}

func (item *feedbackItem) GetAccount() string{
	return item.Account
}
