package main

type ApplianceData struct {
	Id               string `json:"_id"`
	ApplianceId      string `json:"applianceId"`
	LastFullSyncTime int64  `json:"lastFullSyncTime"`
	LastModified     int64  `json:"lastModified"`
	Attributes       struct {
		XcatApplianceinfosetserialnumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_ApplianceInfoSetSerialNumber"`
		CavityOpsethorzlouverswing struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetHorzLouverSwing"`
		SysAlertstatuscustomerfaultcode struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_AlertStatusCustomerFaultCode"`
		SysOpsettargettime struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetTargetTime"`
		SysDisplaysetbrightness struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_DisplaySetBrightness"`
		XCatOnline struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_Online"`
		CavityOpsetresetairfiltertime struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetResetAirFilterTime"`
		SysOpsettargettemp struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetTargetTemp"`
		CavityOpsetmode struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetMode"`
		XcatWifisetdeprovisionwificommmodule struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_WifiSetDeprovisionWifiCommModule"`
		XcatPowerstatusrealtimepower struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerStatusRealTimePower"`
		CavityOpsetvertlouverswing struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetVertLouverSwing"`
		XcatPowerstatusrealtimecurrent struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerStatusRealTimeCurrent"`
		ModelNumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ModelNumber"`
		IspPartNumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ISP_PART_NUMBER"`
		XcatOdometerstatusrunninghours struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_OdometerStatusRunningHours"`
		SysOpsetpoweron struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetPowerOn"`
		ApplianceVersionNumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ApplianceVersionNumber"`
		Version struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"version"`
		Ccuri struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ccuri"`
		CavityOpstatusairfilterstatus struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpStatusAirFilterStatus"`
		XcatPersistentinfoversion struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PersistentInfoVersion"`
		CavityOpsetfanspeed struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetFanSpeed"`
		SerialNumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"SerialNumber"`
		CavityOpsetairfilter struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetAirFilter"`
		XcatPowerstatusrealtimevoltage struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerStatusRealTimeVoltage"`
		XcatWifisetdisablewificommmodule struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_WifiSetDisableWifiCommModule"`
		XcatWifistatusrssiantennadiversity struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_WifiStatusRssiAntennaDiversity"`
		XcatPowerstatuspoweroutage struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerStatusPowerOutage"`
		XcatOdometerstatustotalhours struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_OdometerStatusTotalHours"`
		XcatPersistentinfosaid struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PersistentInfoSaid"`
		SysOpsetecomodeenabled struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetEcoModeEnabled"`
		XcatWifisetpublishappliancestate struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_WifiSetPublishApplianceState"`
		SAID struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"SAID"`
		SysDisplaysettempunits struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_DisplaySetTempUnits"`
		XcatPowersetrealtimepowerpublishtiming struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerSetRealTimePowerPublishTiming"`
		SysOpstatusdisplaytemp struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpStatusDisplayTemp"`
		XcatApplianceinfosetmodelnumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_ApplianceInfoSetModelNumber"`
		SysOpsettargethumidity struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetTargetHumidity"`
		SysOpsetquietmodeenabled struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetQuietModeEnabled"`
		SysOpsetsleepmode struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpSetSleepMode"`
		CavityOpsetturbomode struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetTurboMode"`
		SysOpstatusdisplayhumidity struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Sys_OpStatusDisplayHumidity"`
		XcatWifisetrebootwificommmodule struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_WifiSetRebootWifiCommModule"`
		CavityOpstatusmode struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpStatusMode"`
		IspWifiVersion struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ISP_WIFI_VERSION"`
		XcatPersistentinfomacaddress struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PersistentInfoMacAddress"`
		XcatDatetimesetdatetimeset struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_DateTimeSetDateTimeSet"`
		CavityOpsetairfiltertime struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Cavity_OpSetAirFilterTime"`
		TimeZoneId struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"TimeZoneId"`
		IspWifiPartNumber struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"ISP_WIFI_PART_NUMBER"`
		XcatOdometerstatuscyclecount struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_OdometerStatusCycleCount"`
		Online struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"Online"`
		MacAddress struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"MAC_Address"`
		XcatPowersetrealtimepowerpublishstart struct {
			Value      string `json:"value"`
			UpdateTime int64  `json:"updateTime"`
		} `json:"XCat_PowerSetRealTimePowerPublishStart"`
	} `json:"attributes"`
}
