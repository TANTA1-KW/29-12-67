import { useState, useEffect } from "react";
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
  Upload,
  Select,
  InputNumber,
} from "antd";
import dayjs from "dayjs";
import { PlusOutlined } from "@ant-design/icons";
import { CustomersInterface } from "../../../interfaces/Customer";
import { GendersInterface } from "../../../interfaces/Gender";
import { GetGender, GetUsersById, UpdateUsersById } from "../../../services/https";
import { useNavigate, useParams, Link } from "react-router-dom";
import { UploadFile } from "antd/es/upload/interface";
const { Option } = Select;
import './index.css';

function CustomerEdit() {
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [genders, setGenders] = useState<GendersInterface[]>([]);
  const { id } = useParams<{ id: any }>();
  const [form] = Form.useForm();
  const [fileList, setFileList] = useState<UploadFile[]>([]);

  const onFinish = async (values: CustomersInterface) => {
    let payload = { ...values };
    payload.picture = fileList[0]?.originFileObj
      ? await convertToBase64(fileList[0]?.originFileObj)
      : fileList[0]?.url;

    const res = await UpdateUsersById(id, payload);
    if (res.status === 200) {
      messageApi.open({
        type: "success",
        content: res.data.message,
      });
      setTimeout(() => {
        navigate("/customer");
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: res.data.error,
      });
    }
  };

  const getGender = async () => {
    let res = await GetGender();
    if (res.status) {
      setGenders(res.data);
    }
  };

  const getUserById = async (id: number) => {
    let res = await GetUsersById(id);
    if (res.status === 200) {
      form.setFieldsValue({
        first_name: res.data.first_name,
        last_name: res.data.last_name,
        email: res.data.email,
        birthday: dayjs(res.data.Birthday),
        age: res.data.age,
        gender_id: res.data.gender?.ID,
        phone: res.data.phone,
        address: res.data.Address,
      });
      setFileList(
        res.data.picture
          ? [
              {
                uid: "-1",
                name: "current-image.png",
                status: "done",
                url: res.data.picture,
              },
            ]
          : []
      );
    } else {
      messageApi.open({
        type: "error",
        content: "ไม่พบข้อมูลผู้ใช้",
      });
      setTimeout(() => {
        navigate("/customer");
      }, 2000);
    }
  };

  useEffect(() => {
    getGender();
    if (id) {
      getUserById(id);
    }
  }, [id]);

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

  return (
    <div className="card-container"style={{ height: "150vh", display: "flex", justifyContent: "center", alignItems: "center" }}>
      {contextHolder}
      <Card style={{ width: "100%", maxWidth: 800 }}>
        <h2>แก้ไขข้อมูล ผู้ดูแลระบบ</h2>
        <Divider />
        <Form
          name="basic"
          form={form}
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
                beforeUpload={() => false} // ไม่ให้อัปโหลดไปที่เซิร์ฟเวอร์ทันที
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
                label="ชื่อจริง"
                name="first_name"
                rules={[{ required: true, message: "กรุณากรอก ชื่อ !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item
                label="นามสกุล"
                name="last_name"
                rules={[{ required: true, message: "กรุณากรอก นามสกุล !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item
                name="gender_id"
                label="เพศ"
                rules={[{ required: true, message: "กรุณาระบุ เพศ !" }]}
              >
                <Select allowClear>
                  {genders.map((item) => (
                    <Option value={item.ID} key={item.gender}>
                      {item.gender}
                    </Option>
                  ))}
                </Select>
              </Form.Item>                      
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item label="อายุ" name="age" rules={[{ required: true, message: "กรุณากรอก อายุ !" }]}>
                <InputNumber min={0} max={150} defaultValue={0} style={{ width: "100%" }} />
              </Form.Item>
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item
                label="โทรศัพท์"
                name="phone"
                rules={[
                  { required: true, message: "กรุณากรอก เบอร์โทร !" },
                  { pattern: /^[0]\d{9}$/, message: "Phone number must be exactly 10 digits and start with '0'!" },
                ]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item label="วัน/เดือน/ปี เกิด" name="birthday" rules={[{ required: true, message: "กรุณาเลือกวันเกิด !" }]}>
                <DatePicker style={{ width: "100%" }} />
              </Form.Item>
            </Col>
            <Col xs={24} sm={24} md={24} lg={24} xl={12}>
              <Form.Item label="ที่อยู่" name="address" rules={[{ required: true, message: "กรุณากรอกที่อยู่ !" }]}>
                <Input.TextArea rows={1} />
              </Form.Item>
            </Col>
          </Row>

          <Row justify="end">
            <Col style={{ marginTop: "40px" }}>
              <Form.Item>
                <Space>
                  <Link to="/customer">
                    <Button htmlType="button" style={{ marginRight: "10px" }}>
                      ยกเลิก
                    </Button>
                  </Link>
                  <Button type="primary" htmlType="submit" icon={<PlusOutlined />}>
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

export default CustomerEdit;
