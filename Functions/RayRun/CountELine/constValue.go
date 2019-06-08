package CountELine

const (
	STEP          = 0.5 //电子计算的步长
	START         =95.0
	END           =1000.0
	AM            = "02"
	PM            = "14.50"
	XiAnLatitude  = "34"
	XiAnLongitude = "108"
	IRIName       = "IRI Mode :"
	PROJECTNAME   ="Project Mode :"
	IRIPOSTHead = `
   model=iri2016
   time_flag=1
   height=100.
   geo_flag=0.
   profile=1
   step=0.5
   sun_n=
   ion_n=
   radio_f=
   radio_f81=
   htec_max=
   ne_top=0.
   imap=0.
   ffof2=0.
   hhmf2=0.
   ib0=2.
   probab=0.
   fauroralb=1.
   ffoE=1.
   dreg=0.
   tset=0.
   icomp=0.
   nmf2=0.
   hmf2=0.
   nme=0.
   user_hme=0.
   user_B0=0.
   format=2
   vars=06
   vars=17
   linestyle=solid
   charsize=
   symbol=2
   symsize=
   yscale=Linear
   xscale=Linear
   imagex=640
   imagey=480`
    IRIPOSTYEAR =" year="
    IRIPOSTMONTH=" month="
    IRIPOSTDAY=" day="
    IRIPOSTHOUR=" hour="
    IRIPOSTSTART=" start="
    IRIPOSTSTOP=" stop="
    IRIPOSTLATITUDE=" latitude="
    IRIPOSTLONGITUDE=" longitude="
	IRIWEBSIDE ="https://ccmc.gsfc.nasa.gov/cgi-bin/modelweb/models/vitmo_model.cgi"
	IRIDIRWEBSIDE="https://ccmc.gsfc.nasa.gov/idl_images/"
)
