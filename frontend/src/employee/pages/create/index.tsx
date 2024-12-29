import {
    Space,
    Button,
    Col,
    Row,
    Divider,
    Form,
    Input,
    Card,
    message,
    DatePicker,
    InputNumber,
    Select,
    Upload,
  
  } from "antd";
  import { useState, useEffect } from "react";
  import { PlusOutlined } from "@ant-design/icons";
  import { EmployeesInterface } from "../../interface/Employee";
  import { GendersInterface } from "../../../interfaces/Gender";
  import { RolesInterface } from "../../interface/Role";
  import { StatusInterface } from "../../interface/Status";
  import { GetGender} from "../../../services/https";
  import { GetRole,GetStatus,CreateEmployee } from "../../service/https";
  import { useNavigate, Link } from "react-router-dom";
  import { UploadFile } from "antd/es/upload/interface";
  import './index.css';
  function EmployeeCreate() {
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();
    const [gender, setGender] = useState<GendersInterface[]>([]);
    const [role, setRole] = useState<RolesInterface[]>([]);
    const [stat, setStat] = useState<StatusInterface[]>([]);
    const [fileList, setFileList] = useState<UploadFile[]>([]);

    const onGetGender = async () => {
      let res = await GetGender();
      if (res.status == 200) {
        setGender(res.data);
      } else {
        messageApi.open({
          type: "error",
          content: "ไม่พบข้อมูลเพศ",
        });
        setTimeout(() => {
          navigate("/employee");
        }, 2000);
      }
    };

    const onGetRole = async () => {
        let res = await GetRole();
        if (res.status == 200) {
          setRole(res.data);
        } else {
          messageApi.open({
            type: "error",
            content: "ไม่พบข้อมูลเพศ",
          });
          setTimeout(() => {
            navigate("/employee");
          }, 2000);
        }
      };

      const onGetStat = async () => {
        let res = await GetStatus();
        if (res.status == 200) {
          setStat(res.data);
        } else {
          messageApi.open({
            type: "error",
            content: "ไม่พบข้อมูลเพศ",
          });
          setTimeout(() => {
            navigate("/employee");
          }, 2000);
        }
      };
  
    const onFinish = async (values: EmployeesInterface) => {
      let payload = { ...values };
      payload.picture = fileList[0]?.originFileObj
        ? await convertToBase64(fileList[0]?.originFileObj)
        : fileList[0]?.url;


      const res = await CreateEmployee(payload);
      if (res.status == 201) {
        messageApi.open({
          type: "success",
          className: "message-success",
          content: "สร้างข้อมูลสำเร็จ!",
        });
        setTimeout(() => {
          navigate("/employee");
        }, 2000);
      } else {
        messageApi.open({
          type: "error",
          content: res.data.error || "ไม่สามารถสร้างข้อมูลได้!",
        });
      }
    };

      const convertToBase64 = (file: Blob): Promise<string> => {
        return new Promise((resolve, reject) => {
          const reader = new FileReader();
          reader.onloadend = () => resolve(reader.result as string);
          reader.onerror = reject;
          reader.readAsDataURL(file);
        });
      };
    
      const handleFileChange = ({ fileList }: { fileList: UploadFile[] }) => {
        setFileList(fileList);
      };
  
    useEffect(() => {
      onGetGender();
      onGetRole();
      onGetStat();
      return () => {};
    }, []);
  
    return (
      <div style={{ height: "150vh", display: "flex", justifyContent: "center", alignItems: "center" }}>
        {contextHolder}
        <Card style={{ width: "100%", maxWidth: 800 }}>
          <h2>สร้าง Account Employee</h2>
          <Divider />
          <Form
            name="basic"
            layout="vertical"
            onFinish={onFinish}
            autoComplete="off"
          >
          <Row justify="center">
            <Form.Item name="picture" label="" className="upload-image">
              <Upload
                listType="picture-card"
                fileList={fileList}
                onChange={handleFileChange}
                beforeUpload={(file) => {
                  setFileList([...fileList, file]);
                  return false;
                }} // ไม่ให้อัปโหลดไปที่เซิร์ฟเวอร์ทันที
                showUploadList={{ showPreviewIcon: false }}
                maxCount={1}
              >
                {fileList.length < 1 && (
                    <div>
                      <PlusOutlined />
                    <div style={{ marginTop: 8 }}>อัปโหลด</div>
                    </div>
                  )}
                </Upload>
              </Form.Item>
            </Row>
            <Row gutter={[16, 16]}>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="ตำแหน่ง"
                  name="role_id"
                  rules={[
                    {
                      required: true,
                      message: "กรุณาเลือกตำแหน่ง !",
                    },
                  ]}
                >
                   <Select allowClear>
                    {role?.map((item) => (
                      <Select.Option
                        value={item?.ID} key={item.role}
                      >
                        {item?.role}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="เพศ"
                  name="gender_id"
                  rules={[
                    {
                      required: true,
                      message: "กรุณาเลือกเพศ !",
                    },
                  ]}
                >
                   <Select allowClear>
                    {gender?.map((item) => (
                      <Select.Option
                        value={item?.ID} key={item.gender}
                      >
                        {item?.gender}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>

              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="ชื่อจริง"
                  name="first_name"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกชื่อ !",
                    },
                  ]}
                >
                  <Input />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="นามสกุล"
                  name="last_name"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกนามสกุล !",
                    },
                  ]}
                >
                  <Input />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="อีเมล"
                  name="email"
                  rules={[
                    {
                      type: "email",
                      message: "รูปแบบอีเมลไม่ถูกต้อง !",
                    },
                    {
                      required: true,
                      message: "กรุณากรอกอีเมล !",
                    },
                  ]}
                >
                  <Input />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="รหัสผ่าน"
                  name="password"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกรหัสผ่าน !",
                    },
                  ]}
                >
                  <Input.Password />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="วัน/เดือน/ปี เกิด"
                  name="birthday"
                  rules={[
                    {
                      required: true,
                      message: "กรุณาเลือกวัน/เดือน/ปี เกิด !",
                    },
                  ]}
                >
                  <DatePicker style={{ width: "100%" }} />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="อายุ"
                  name="age"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกอายุ !",
                    },
                  ]}
                >
                  <InputNumber
                    min={0}
                    max={99}
                    defaultValue={0}
                    style={{ width: "100%" }}
                  />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item
                label="เบอร์โทรศัพท์"
                name="Phone"
                rules={[
                  {
                    required: true,
                    message: "กรุณากรอกเบอร์โทรศัพท์ !",
                  },
                  {
                    pattern: /^[0]\d{9}$/,
                    message: "กรุณากรอก เบอร์โทร",
                  },
                ]}
              >
                <Input />
              </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="เงินเดือน"
                  name="salary"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกเงินเดือน!",
                    },
                  ]}
                >
                  <InputNumber
                    min={0}
                    max={99999}
                    defaultValue={0}
                    style={{ width: "100%" }}
                    precision={2}
                  />
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="สถานะ"
                  name="stat_id"
                  rules={[
                    {
                      required: true,
                      message: "กรุณาเลือกสถานะ !",
                    },
                  ]}
                >
                   <Select allowClear>
                    {stat?.map((item) => (
                      <Select.Option
                        value={item?.ID} key={item?.stat}
                      >
                        {item?.stat}
                      </Select.Option>
                    ))}
                  </Select>
                </Form.Item>
              </Col>
              <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                <Form.Item
                  label="ที่อยู่"
                  name="address"
                  rules={[
                    {
                      required: true,
                      message: "กรุณากรอกที่อยู่ !",
                    },
                  ]}
                >
                  <Input.TextArea rows={1} />
                </Form.Item>
              </Col>
            </Row>
            <Row justify="end">
              <Col style={{ marginTop: "40px" }}>
                <Form.Item>
                  <Space>
                    <Link to="/employee">
                      <Button htmlType="button" style={{ marginRight: "10px" }}>
                        ยกเลิก
                      </Button>
                    </Link>
                    <Button
                      type="primary"
                      htmlType="submit"
                      icon={<PlusOutlined />}
                    >
                      ยืนยัน
                    </Button>
                  </Space>
                </Form.Item>
              </Col>
            </Row>
          </Form>
        </Card>
      </div>
    );
  }
  
export default EmployeeCreate;