import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Textarea } from "./ui/textarea";
import { Button } from "./ui/button";
import { Popover, PopoverContent, PopoverTrigger } from "./ui/popover";
import { Check, ChevronsUpDown } from "lucide-react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "./ui/command";
import { cn } from "@/lib/utils";
import { FC } from "react";
import {
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "./ui/dialog";
import { useToast } from "./ui/use-toast";

const formSchema = z.object({
  question: z
    .string({
      message: "Вопрос должен быть больше 4 символов",
    })
    .min(4),
  answer: z
    .string({
      message: "Ответ должен быть больше 4 символов",
    })
    .min(4),
  sub_category_id: z
    .number({
      message: "Выберите подтему",
    })
    .min(1),
});

interface SubCategory {
  label: string;
  value: number;
}

interface AddCardProps {
  subCategories: SubCategory[];
}

export const AddCard: FC<AddCardProps> = ({ subCategories }) => {
  const { toast } = useToast();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      question: "",
      answer: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const res = await fetch("http://localhost:5000/api/card/create", {
      method: "POST",
      body: JSON.stringify(values),
    });

    const data = await res.json();

    toast({
      title: "Ответ от сервера",
      description: (
        <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
          <code className="text-white">{JSON.stringify(data, null, 2)}</code>
        </pre>
      ),
    });
  };

  return (
    <DialogContent className="max-w-3xl">
      <DialogHeader>
        <DialogTitle>Создание карточки</DialogTitle>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <div className="px-4 py-4 sm:px-0">
            <FormField
              control={form.control}
              name="sub_category_id"
              render={({ field }) => (
                <FormItem className="flex flex-col">
                  <FormLabel>Подтемы</FormLabel>
                  <Popover>
                    <PopoverTrigger asChild>
                      <FormControl>
                        <Button
                          variant="outline"
                          role="combobox"
                          className={cn(
                            "w-full justify-between",
                            !field.value && "text-muted-foreground",
                          )}
                        >
                          {field.value
                            ? subCategories.find(
                              (subCategory) =>
                                subCategory.value === field.value,
                            )?.label
                            : "Выберите подтему"}
                          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                        </Button>
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent className="w-full p-0">
                      <Command>
                        <CommandInput placeholder="Поиск ..." />
                        <CommandList>
                          <CommandEmpty>Пусто.</CommandEmpty>
                          <CommandGroup>
                            {subCategories.map((subCategory) => (
                              <CommandItem
                                value={subCategory.label}
                                key={subCategory.value}
                                onSelect={() => {
                                  form.setValue(
                                    "sub_category_id",
                                    subCategory.value,
                                  );
                                }}
                              >
                                <Check
                                  className={cn(
                                    "mr-2 h-4 w-4",
                                    subCategory.value === field.value
                                      ? "opacity-100"
                                      : "opacity-0",
                                  )}
                                />
                                {subCategory.label}
                              </CommandItem>
                            ))}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className="divide-y divide-gray-200">
            <FormField
              control={form.control}
              name="question"
              render={({ field }) => (
                <FormItem className="px-4 py-4 sm:px-0">
                  <FormLabel>Вопрос и Ответ</FormLabel>
                  <FormControl>
                    <Textarea placeholder="Напишите вопрос" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="answer"
              render={({ field }) => (
                <FormItem className="px-4 py-4 sm:px-0">
                  <FormControl>
                    <Textarea
                      placeholder="Напишите ответ к вопросу"
                      {...field}
                    />
                  </FormControl>
                  <FormDescription>
                    Для обозначения формул напишите $ перед и после формулы,
                    если нужно большую формулу обозначить, то используйте $$
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <DialogFooter>
            <Button className="w-full" type="submit">
              Создать
            </Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};
