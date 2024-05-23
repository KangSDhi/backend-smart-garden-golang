package dto

type RequestSoilMoistureSensor struct {
	NilaiSensorAnalog uint   `json:"nilai_sensor_analog" validate:"required"`
	LabelSensorAnalog string `json:"label_sensor_analog" validate:"required"`
	TanggalSensor     string `json:"tanggal_sensor" validate:"required,datetime=2006-01-02 15:04:05"`
}
