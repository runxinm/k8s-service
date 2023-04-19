// func CreateNamespace(){
	// 循环50次创建namespace, 使用 goroutine 并发创建 但是要注意加锁，否则同名namespaces已经存在时，会出现错误
		// wg.Add(1) means add 1 goroutine
		// wg.Done() means one goroutine is done
		// wg.Wait() means wait all goroutine is done
		// the goroutine is not a thread, it is a function
		// Add Wait Done 顺序是Add Done Wait
	// var wg sync.WaitGroup
	// wg.Add(50)
	// for i := 0; i < 50; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		k8sns.Name = fmt.Sprintf("test-%d",i)
	// 		fmt.Println(k8sns.Name)
	// 		_, err := clientset.CoreV1().Namespaces().Create(context.Background(),k8sns,metav1.CreateOptions{})
	// 		// _, err = clientset.CoreV1().Namespaces().Create(k8sns)
	// 		if err != nil {
	// 			fmt.Println("create namespace error")
	// 		}
	// 	}()
	// }


	// wg := sync.WaitGroup{}
	// var mtx sync.RWMutex
	// var wg sync.WaitGroup
	// wg.Add(50)
	// for i := 0; i < 50; i++ {
		// go func() {
			// defer wg.Done()
			// mtx.Lock()
	// 		tmpns := &corev1.Namespace{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Name: "1",
	// 			},
	// 		}
	// 		tmpns.Name = fmt.Sprintf("test-%d",i)
	// 		fmt.Println(tmpns.Name)
	// 		_, err := clientset.CoreV1().Namespaces().Create(context.Background(),tmpns,metav1.CreateOptions{})
	// 		// mtx.Unlock()
	// 		// _, err = clientset.CoreV1().Namespaces().Create(k8sns)
	// 		if err != nil {
	// 			fmt.Println("create namespace error")
	// 		}
	// 	// }()
	// }
	// // wg.Wait()
	// // sleep 5s
	// fmt.Println("sleep 5s")
	// time.Sleep(5 * time.Second)
// }