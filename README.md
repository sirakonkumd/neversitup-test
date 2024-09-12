# neversitup-test

structure project ที่ใช้ปัจจุบันจะใช้ pattern Hexagonal Architecture
ซึ่ง core หลักจะมี service ที่ใช้ทำ business logic รวมไปถึงการเรีกใช้ repositories, thirdparties และอื่นๆ
routes จะรวมการทำ server, setup route, handler
ใน routes/route.go จะใช้ทำ setup server
ใน routes/handlers/register.go จะใช้ setup path api
ใน routes/handlers/serverInterface.go เรียก function services เพื่อที่จะเอาค่าจาก struct interface ไปใช้ใน handler
ส่วนอื่นๆใน handler จะเป็นตัวทำการเรียกใช้ service ซึ่งจะแบ่งการเรียกใช้ service ตามชื่อไฟล์ของ handler

thirdparties จะเป็นส่วนในการ call api service อื่นๆ
ีutils จะเก็บพวก func การคำนวนทั่วไปเพื่อให้ service ต่างๆได้ใช้งาน เช่น ParseToTime, formatDate , containSliceString เป็นต้น
repositories จะเป็นในส่วนของการ query ข้อมูลจาก database
ตัวอย่างที่ service เรียกใช้ repositories จะอยู่ที่ services/oddServices ครับ
models จะเก็บ struct ที่ใช้ในระดับ global
database จะเป็นการทำ connection ของแต่ละ db
constants จะเก็บค่า constant ที่จะใช้กับ service อื่นๆ
